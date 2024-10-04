package eds

import (
	"context"
	"io"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/require"

	"github.com/celestiaorg/go-square/v2/share"
	"github.com/celestiaorg/rsmt2d"

	"github.com/celestiaorg/celestia-node/square"
	"github.com/celestiaorg/celestia-node/square/shwap"
)

func TestWithClosedOnce(t *testing.T) {
	ctx := context.Background()
	stub := &stubEdsAccessorCloser{}
	closedOnce := WithClosedOnce(stub)

	_, err := closedOnce.Sample(ctx, 0, 0)
	require.NoError(t, err)
	_, err = closedOnce.AxisHalf(ctx, rsmt2d.Row, 0)
	require.NoError(t, err)
	_, err = closedOnce.RowNamespaceData(ctx, share.Namespace{}, 0)
	require.NoError(t, err)
	_, err = closedOnce.Shares(ctx)
	require.NoError(t, err)

	require.NoError(t, closedOnce.Close())
	require.True(t, stub.closed)

	// Ensure that the underlying file is not accessible after closing
	_, err = closedOnce.Sample(ctx, 0, 0)
	require.ErrorIs(t, err, errAccessorClosed)
	_, err = closedOnce.AxisHalf(ctx, rsmt2d.Row, 0)
	require.ErrorIs(t, err, errAccessorClosed)
	_, err = closedOnce.RowNamespaceData(ctx, share.Namespace{}, 0)
	require.ErrorIs(t, err, errAccessorClosed)
	_, err = closedOnce.Shares(ctx)
	require.ErrorIs(t, err, errAccessorClosed)
}

type stubEdsAccessorCloser struct {
	closed bool
}

func (s *stubEdsAccessorCloser) Size(context.Context) int {
	return 0
}

func (s *stubEdsAccessorCloser) DataHash(context.Context) (square.DataHash, error) {
	return square.DataHash{}, nil
}

func (s *stubEdsAccessorCloser) AxisRoots(context.Context) (*square.AxisRoots, error) {
	return &square.AxisRoots{}, nil
}

func (s *stubEdsAccessorCloser) Sample(context.Context, int, int) (shwap.Sample, error) {
	return shwap.Sample{}, nil
}

func (s *stubEdsAccessorCloser) AxisHalf(context.Context, rsmt2d.Axis, int) (AxisHalf, error) {
	return AxisHalf{}, nil
}

func (s *stubEdsAccessorCloser) RowNamespaceData(
	context.Context,
	share.Namespace,
	int,
) (shwap.RowNamespaceData, error) {
	return shwap.RowNamespaceData{}, nil
}

func (s *stubEdsAccessorCloser) Shares(context.Context) ([]share.Share, error) {
	return nil, nil
}

func (s *stubEdsAccessorCloser) Reader() (io.Reader, error) {
	return iotest.ErrReader(nil), nil
}

func (s *stubEdsAccessorCloser) Close() error {
	s.closed = true
	return nil
}