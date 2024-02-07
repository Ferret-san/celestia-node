package store

import (
	"context"
	"errors"
	"fmt"

	bstore "github.com/ipfs/boxo/blockstore"
	"github.com/ipfs/boxo/datastore/dshelp"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/namespace"
	ipld "github.com/ipfs/go-ipld-format"

	"github.com/celestiaorg/celestia-node/share/shwap"
)

var _ bstore.Blockstore = (*Blockstore)(nil)

var (
	blockstoreCacheKey      = datastore.NewKey("bs-cache")
	errUnsupportedOperation = errors.New("unsupported operation")
)

// Blockstore implements the bstore.Blockstore interface on an EDSStore.
// It is used to provide a custom blockstore interface implementation to achieve access to the
// underlying EDSStore. The main use-case is randomized sampling over the whole chain of EDS block
// data and getting data by namespace.
type Blockstore struct {
	store *Store
	ds    datastore.Batching
}

func NewBlockstore(store *Store, ds datastore.Batching) *Blockstore {
	return &Blockstore{
		store: store,
		ds:    namespace.Wrap(ds, blockstoreCacheKey),
	}
}

func (bs *Blockstore) Has(ctx context.Context, cid cid.Cid) (bool, error) {
	h, err := shwap.BlockBuilderFromCID(cid)
	if err != nil {
		return false, fmt.Errorf("while getting height from CID: %w", err)
	}

	// check cache first
	height := h.GetHeight()
	_, err = bs.store.cache.Get(height)
	if err == nil {
		return true, nil
	}

	_, err = bs.store.HasByHeight(ctx, height)
	if err == nil {
		return true, nil
	}
	if !errors.Is(err, ErrNotFound) {
		return false, fmt.Errorf("failed to get file: %w", err)
	}

	// key wasn't found in top level blockstore, but could be in datastore while being reconstructed
	dsHas, dsErr := bs.ds.Has(ctx, dshelp.MultihashToDsKey(cid.Hash()))
	if dsErr != nil {
		return false, nil
	}
	return dsHas, nil
}

func (bs *Blockstore) Get(ctx context.Context, cid cid.Cid) (blocks.Block, error) {
	h, err := shwap.BlockBuilderFromCID(cid)
	if err != nil {
		return nil, fmt.Errorf("while getting height from CID: %w", err)
	}

	height := h.GetHeight()
	f, err := bs.store.cache.Second().GetOrLoad(ctx, height, bs.store.openFileByHeight(height))
	if err == nil {
		return h.BlockFromFile(ctx, f)
	}

	if errors.Is(err, ErrNotFound) {
		k := dshelp.MultihashToDsKey(cid.Hash())
		blockData, err := bs.ds.Get(ctx, k)
		if err == nil {
			return blocks.NewBlockWithCid(blockData, cid)
		}
		// nmt's GetNode expects an ipld.ErrNotFound when a cid is not found.
		return nil, ipld.ErrNotFound{Cid: cid}
	}

	log.Debugf("failed to get blockstore for cid %s: %s", cid, err)
	return nil, err
}

func (bs *Blockstore) GetSize(ctx context.Context, cid cid.Cid) (int, error) {
	// TODO(@Wondertan): There must be a way to derive size without reading, proving, serializing and
	//  allocating Sample's block.Block.
	// NOTE:Bitswap uses GetSize also to determine if we have content stored or not
	// so simply returning constant size is not an option
	blk, err := bs.Get(ctx, cid)
	if err != nil {
		return 0, err
	}

	return len(blk.RawData()), nil
}

func (bs *Blockstore) DeleteBlock(ctx context.Context, cid cid.Cid) error {
	k := dshelp.MultihashToDsKey(cid.Hash())
	return bs.ds.Delete(ctx, k)
}

func (bs *Blockstore) Put(ctx context.Context, blk blocks.Block) error {
	k := dshelp.MultihashToDsKey(blk.Cid().Hash())
	// note: we leave duplicate resolution to the underlying datastore
	return bs.ds.Put(ctx, k, blk.RawData())
}

func (bs *Blockstore) PutMany(ctx context.Context, blocks []blocks.Block) error {
	if len(blocks) == 1 {
		// performance fast-path
		return bs.Put(ctx, blocks[0])
	}

	t, err := bs.ds.Batch(ctx)
	if err != nil {
		return err
	}
	for _, b := range blocks {
		k := dshelp.MultihashToDsKey(b.Cid().Hash())
		err = t.Put(ctx, k, b.RawData())
		if err != nil {
			return err
		}
	}
	return t.Commit(ctx)
}

// AllKeysChan is a noop on the EDS blockstore because the keys are not stored in a single CAR file.
func (bs *Blockstore) AllKeysChan(context.Context) (<-chan cid.Cid, error) {
	err := fmt.Errorf("AllKeysChan is: %w", errUnsupportedOperation)
	log.Warn(err)
	return nil, err
}

// HashOnRead is a noop on the EDS blockstore but an error cannot be returned due to the method
// signature from the blockstore interface.
func (bs *Blockstore) HashOnRead(bool) {
	log.Warn("HashOnRead is a noop on the EDS blockstore")
}