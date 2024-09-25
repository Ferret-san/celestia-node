package file

import (
	"fmt"
	"io"

	"golang.org/x/sync/errgroup"

	"github.com/celestiaorg/go-square/v2/share"
	"github.com/celestiaorg/rsmt2d"

	"github.com/celestiaorg/celestia-node/square/eds"
)

type dataSquare [][]share.Share

// readSquare reads Shares from the reader and returns a square. It assumes that the reader is
// positioned at the beginning of the Shares. It knows the size of the Shares and the size of the
// square, so reads from reader are limited to exactly the amount of data required.
func readSquare(r io.Reader, shareSize, edsSize int) (dataSquare, error) {
	odsLn := edsSize / 2

	shares, err := eds.ReadShares(r, shareSize, odsLn)
	if err != nil {
		return nil, fmt.Errorf("reading shares: %w", err)
	}
	square := make(dataSquare, odsLn)
	for i := range square {
		square[i] = shares[i*odsLn : (i+1)*odsLn]
	}
	return square, nil
}

func (s dataSquare) reader() (io.Reader, error) {
	if s == nil {
		return nil, fmt.Errorf("ods file not cached")
	}
	getShare := func(rowIdx, colIdx int) ([]byte, error) {
		return s[rowIdx][colIdx].ToBytes(), nil
	}
	reader := eds.NewShareReader(s.size(), getShare)
	return reader, nil
}

func (s dataSquare) size() int {
	return len(s)
}

func (s dataSquare) shares() ([]share.Share, error) {
	shares := make([]share.Share, 0, s.size()*s.size())
	for _, row := range s {
		shares = append(shares, row...)
	}
	return shares, nil
}

func (s dataSquare) axisHalf(axisType rsmt2d.Axis, axisIdx int) (eds.AxisHalf, error) {
	if s == nil {
		return eds.AxisHalf{}, fmt.Errorf("square is nil")
	}

	if axisIdx >= s.size() {
		return eds.AxisHalf{}, fmt.Errorf("index is out of square bounds")
	}

	// square stores rows directly in high level slice, so we can return by accessing row by index
	if axisType == rsmt2d.Row {
		row := s[axisIdx]
		return eds.AxisHalf{
			Shares:   row,
			IsParity: false,
		}, nil
	}

	// construct half column from row ordered square
	col := make([]share.Share, s.size())
	for i := 0; i < s.size(); i++ {
		col[i] = s[i][axisIdx]
	}
	return eds.AxisHalf{
		Shares:   col,
		IsParity: false,
	}, nil
}

func (s dataSquare) computeAxisHalf(
	axisType rsmt2d.Axis,
	axisIdx int,
) (eds.AxisHalf, error) {
	shares := make([]share.Share, s.size())

	// extend opposite half of the square while collecting Shares for the first half of required axis
	g := errgroup.Group{}
	opposite := oppositeAxis(axisType)
	for i := 0; i < s.size(); i++ {
		g.Go(func() error {
			half, err := s.axisHalf(opposite, i)
			if err != nil {
				return err
			}

			enc, err := codec.Encoder(s.size() * 2)
			if err != nil {
				return fmt.Errorf("getting encoder: %w", err)
			}

			shards := make([][]byte, s.size()*2)
			if half.IsParity {
				copy(shards[s.size():], share.ToBytes(half.Shares))
			} else {
				copy(shards, share.ToBytes(half.Shares))
			}

			target := make([]bool, s.size()*2)
			target[axisIdx] = true

			err = enc.ReconstructSome(shards, target)
			if err != nil {
				return fmt.Errorf("reconstruct some: %w", err)
			}

			shard, err := share.NewShare(shards[axisIdx])
			if err != nil {
				return fmt.Errorf("creating share: %w", err)
			}
			shares[i] = *shard
			return nil
		})
	}

	err := g.Wait()
	return eds.AxisHalf{
		Shares:   shares,
		IsParity: false,
	}, err
}

func oppositeAxis(axis rsmt2d.Axis) rsmt2d.Axis {
	if axis == rsmt2d.Col {
		return rsmt2d.Row
	}
	return rsmt2d.Col
}
