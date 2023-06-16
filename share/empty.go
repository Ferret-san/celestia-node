package share

import (
	"bytes"
	"fmt"

	"github.com/celestiaorg/celestia-app/pkg/appconsts"
	"github.com/celestiaorg/celestia-app/pkg/da"
	"github.com/celestiaorg/celestia-app/pkg/shares"
	"github.com/celestiaorg/rsmt2d"
)

// EmptyRoot returns Root of the empty block EDS.
func EmptyRoot() *Root {
	return emptyBlockRoot
}

// EmptyExtendedDataSquare returns the EDS of the empty block data square.
func EmptyExtendedDataSquare() *rsmt2d.ExtendedDataSquare {
	return emptyBlockEDS
}

// EmptyBlockShares returns the shares of the empty block.
func EmptyBlockShares() []Share {
	return emptyBlockShares
}

var (
	emptyBlockRoot   *Root
	emptyBlockEDS    *rsmt2d.ExtendedDataSquare
	emptyBlockShares []Share
)

func init() {
	// compute empty block EDS and DAH for it
	result := shares.TailPaddingShares(appconsts.MinShareCount)
	emptyBlockShares = shares.ToBytes(result)

	eds, err := da.ExtendShares(emptyBlockShares)
	if err != nil {
		panic(fmt.Errorf("failed to create empty EDS: %w", err))
	}
	emptyBlockEDS = eds

	dah := da.NewDataAvailabilityHeader(eds)
	minDAH := da.MinDataAvailabilityHeader()
	if !bytes.Equal(minDAH.Hash(), dah.Hash()) {
		panic(fmt.Sprintf("mismatch in calculated minimum DAH and minimum DAH from celestia-app, "+
			"expected %s, got %s", minDAH.String(), dah.String()))
	}
	emptyBlockRoot = &dah

	// precompute Hash, so it's cached internally to avoid potential races
	emptyBlockRoot.Hash()
}
