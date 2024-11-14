package testnet

import (
	"context"
	"fmt"
	"strings"

	"github.com/celestiaorg/celestia-app/v3/test/e2e/testnet"
	"github.com/celestiaorg/celestia-node/nodebuilder/node"
	"github.com/celestiaorg/celestia-node/test/e2e/prometheus"
	"github.com/celestiaorg/knuu/pkg/instance"
	"github.com/celestiaorg/knuu/pkg/sidecars/observability"
)

type Node struct {
	Name         string
	Type         node.Type
	Version      string
	Instance     *instance.Instance
	rpcProxyHost string
}

type JSONRPCError struct {
	Code    int
	Message string
	Data    string
}

func (e *JSONRPCError) Error() string {
	return fmt.Sprintf("JSONRPC Error - Code: %d, Message: %s, Data: %s", e.Code, e.Message, e.Data)
}

func (nt *NodeTestnet) initInstance(ctx context.Context, opts InstanceOptions) (*instance.Instance, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	ins, err := nt.NewInstance(opts.InstanceName)
	if err != nil {
		return nil, err
	}

	if err := ins.Build().SetImage(ctx, fmt.Sprintf("%s:%s", dockerSrcURL, opts.Version)); err != nil {
		return nil, err
	}

	for _, port := range []int{p2pPort, rpcPort} {
		if err := ins.Network().AddPortTCP(port); err != nil {
			return nil, err
		}
	}

	err = ins.Build().ExecuteCommand("celestia", strings.ToLower(opts.NodeType.String()), "init", "--node.store", remoteRootDir)
	if err != nil {
		return nil, err
	}

	if err := ins.Build().Commit(ctx); err != nil {
		return nil, err
	}

	if opts.consensus == nil {
		opts.SetConsensus(nt.Testnet.Node(0).Instance)
	}

	chainID, err := opts.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	genesisHash, err := opts.GenesisHash(ctx)
	if err != nil {
		return nil, err
	}

	err = ins.Build().SetEnvironmentVariable(celestiaCustomEnv, fmt.Sprintf("%s:%s", chainID, genesisHash))
	if err != nil {
		return nil, err
	}

	obsy, err := nt.createObservability(ctx)
	if err != nil {
		return nil, err
	}

	if err := ins.Sidecars().Add(ctx, obsy); err != nil {
		return nil, err
	}

	// Expose the prometheus exporter on the otel collector instance
	if err := obsy.Instance().Network().AddPortTCP(prometheusExporterPort); err != nil {
		return nil, err
	}

	err = nt.Prometheus.AddScrapeJob(ctx, prometheus.ScrapeJob{
		Name:     opts.InstanceName,
		Target:   fmt.Sprintf("%s:%d", opts.InstanceName, prometheusExporterPort),
		Interval: prometheusScrapeInterval,
	})
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (nt *NodeTestnet) CreateNode(ctx context.Context, opts InstanceOptions, trustedNode *Node) (*Node, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	if opts.executor == nil {
		opts.SetExecutor(nt.executor)
	}

	nodeInst, err := nt.initInstance(ctx, opts)
	if err != nil {
		return nil, err
	}

	//TODO: implement an IsEmpty method for Resources in the app testnet pkg
	if opts.Resources == (testnet.Resources{}) {
		opts.Resources = DefaultBridgeResources
	}

	err = nodeInst.Resources().SetMemory(opts.Resources.MemoryRequest, opts.Resources.MemoryLimit)
	if err != nil {
		return nil, err
	}

	if err := nodeInst.Resources().SetCPU(opts.Resources.CPU); err != nil {
		return nil, err
	}

	startCmd := []string{
		"celestia",
		strings.ToLower(opts.NodeType.String()),
		"start",
		"--node.store", remoteRootDir,
		"--metrics",
		"--metrics.endpoint", fmt.Sprintf("localhost:%d", otlpPort),
		"--metrics.tls=false",
	}

	if opts.NodeType == node.Bridge {
		if opts.consensus == nil {
			opts.SetConsensus(nt.Testnet.Node(0).Instance)
		}
		consensusIP, err := opts.consensus.Network().GetIP(ctx)
		if err != nil {
			return nil, err
		}
		startCmd = append(startCmd, "--core.ip", consensusIP, "--rpc.addr", "0.0.0.0")

	} else {
		trustedPeers, err := getTrustedPeers(ctx, trustedNode)
		if err != nil {
			return nil, err
		}
		startCmd = append(startCmd, "--headers.trusted-peers", trustedPeers)
	}

	if err := nodeInst.Build().SetStartCommand(startCmd...); err != nil {
		return nil, err
	}

	return &Node{
		Name:     opts.InstanceName,
		Type:     opts.NodeType,
		Version:  opts.Version,
		Instance: nodeInst,
	}, nil
}

func (nt *NodeTestnet) CreateAndStartNode(ctx context.Context, opts InstanceOptions, trustedNode *Node) (*Node, error) {
	node, err := nt.CreateNode(ctx, opts, trustedNode)
	if err != nil {
		return nil, ErrFailedToCreateNode.Wrap(err)
	}

	if err := node.Instance.Execution().Start(ctx); err != nil {
		return nil, err
	}

	rpcProxyHost, err := node.Instance.Network().AddHost(ctx, rpcPort)
	if err != nil {
		return nil, err
	}
	node.rpcProxyHost = rpcProxyHost

	return node, nil
}

func (nt *NodeTestnet) createObservability(ctx context.Context) (*observability.Obsy, error) {
	obsy := observability.New()
	if err := obsy.SetOtelEndpoint(otlpPort); err != nil {
		return nil, err
	}

	if err := obsy.SetPrometheusExporter(fmt.Sprintf("0.0.0.0:%d", prometheusExporterPort)); err != nil {
		return nil, err
	}

	return obsy, nil
}

// AddressRPC returns an RPC endpoint address for the node.
// This returns the proxy host that can be used to communicate with the node
func (n Node) AddressRPC() string {
	return n.rpcProxyHost
}

func (n Node) GetNodeID(ctx context.Context) (string, error) {
	out, err := n.Instance.Execution().ExecuteCommand(ctx, "celestia", "p2p", "info | jq -r '.result.id'")
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}
