package share

import (
	"context"
	"fmt"

	"github.com/ipfs/boxo/blockstore"
	"github.com/ipfs/boxo/exchange"
	"github.com/ipfs/go-datastore"
	ipfsmetrics "github.com/ipfs/go-metrics-interface"
	ipfsprom "github.com/ipfs/go-metrics-prometheus"
	hst "github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/fx"

	"github.com/celestiaorg/celestia-node/nodebuilder/node"
	"github.com/celestiaorg/celestia-node/nodebuilder/p2p"
	"github.com/celestiaorg/celestia-node/share/shwap/p2p/bitswap"
	"github.com/celestiaorg/celestia-node/store"
)

// dataExchange constructs Exchange(Bitswap Composition) for Shwap
func dataExchange(tp node.Type, params bitSwapParams) exchange.SessionExchange {
	prefix := protocolID(params.Net)
	net := bitswap.NewNetwork(params.Host, prefix+"_load_test")

	if params.PromReg != nil {
		// metrics scope is required for prometheus metrics and will be used as metrics name prefix
		params.Ctx = ipfsmetrics.CtxScope(params.Ctx, "bitswap")
		err := ipfsprom.Inject()
		if err != nil {
			return nil
		}
	}

	switch tp {
	case node.Full, node.Bridge:
		bs := bitswap.New(params.Ctx, net, params.Bs)
		net.Start(bs.Client, bs.Server)
		params.Lifecycle.Append(fx.Hook{
			OnStop: func(_ context.Context) (err error) {
				net.Stop()
				return bs.Close()
			},
		})
		return bs
	case node.Light:
		cl := bitswap.NewClient(params.Ctx, net, params.Bs)
		net.Start(cl)
		params.Lifecycle.Append(fx.Hook{
			OnStop: func(_ context.Context) (err error) {
				net.Stop()
				return cl.Close()
			},
		})
		return cl
	default:
		panic(fmt.Sprintf("unsupported node type: %v", tp))
	}
}

func blockstoreFromDatastore(ds datastore.Batching) (blockstore.Blockstore, error) {
	return blockstore.NewBlockstore(ds), nil
}

func blockstoreFromEDSStore(store *store.Store, blockStoreCacheSize int) (*bitswap.BlockstoreWithMetrics, error) {
	bs := &bitswap.Blockstore{Getter: store}
	//if blockStoreCacheSize > 0 {
	//	withCache, err := store.WithCache("blockstore", blockStoreCacheSize)
	//	if err != nil {
	//		return nil, fmt.Errorf("create cached store for blockstore:%w", err)
	//	}
	//	bs.Getter = withCache
	//}

	withMetrics, err := bitswap.NewBlockstoreWithMetrics(bs)
	if err != nil {
		return nil, fmt.Errorf("create metrics for blockstore:%w", err)
	}
	return withMetrics, nil
}

type bitSwapParams struct {
	fx.In

	Lifecycle fx.Lifecycle
	Ctx       context.Context
	Net       p2p.Network
	Host      hst.Host
	Bs        blockstore.Blockstore
	PromReg   prometheus.Registerer `optional:"true"`
}

func protocolID(network p2p.Network) protocol.ID {
	return protocol.ID(fmt.Sprintf("/celestia/%s", network))
}
