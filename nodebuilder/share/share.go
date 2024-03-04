package share

import (
	"context"

	"github.com/celestiaorg/rsmt2d"

	"github.com/celestiaorg/celestia-node/header"
	"github.com/celestiaorg/celestia-node/share"
)

var _ Module = (*API)(nil)

// Module provides access to any data square or block share on the network.
//
// All Get methods provided on Module follow the following flow:
//  1. Check local storage for the requested Share.
//  2. If exists
//     * Load from disk
//     * Return
//  3. If not
//     * Find provider on the network
//     * Fetch the Share from the provider
//     * Store the Share
//     * Return
//
// Any method signature changed here needs to also be changed in the API struct.
//
//go:generate mockgen -destination=mocks/api.go -package=mocks . Module
type Module interface {
	// SharesAvailable subjectively validates if Shares committed to the given
	// ExtendedHeader are available on the Network.
	SharesAvailable(context.Context, *header.ExtendedHeader) error
	// GetShare gets a Share by coordinates in EDS.
	GetShare(ctx context.Context, header *header.ExtendedHeader, row, col int) (share.Share, error)
	// GetEDS gets the full EDS identified by the given extended header.
	GetEDS(ctx context.Context, header *header.ExtendedHeader) (*rsmt2d.ExtendedDataSquare, error)
	// GetSharesByNamespace gets all shares from an EDS within the given namespace.
	// Shares are returned in a row-by-row order if the namespace spans multiple rows.
	GetSharesByNamespace(
		ctx context.Context, header *header.ExtendedHeader, namespace share.Namespace,
	) (share.NamespacedShares, error)
}

// API is a wrapper around Module for the RPC.
// TODO(@distractedm1nd): These structs need to be autogenerated.
type API struct {
	Internal struct {
		SharesAvailable func(context.Context, *header.ExtendedHeader) error `perm:"read"`
		GetShare        func(
			ctx context.Context,
			header *header.ExtendedHeader,
			row, col int,
		) (share.Share, error) `perm:"read"`
		GetEDS func(
			ctx context.Context,
			header *header.ExtendedHeader,
		) (*rsmt2d.ExtendedDataSquare, error) `perm:"read"`
		GetSharesByNamespace func(
			ctx context.Context,
			header *header.ExtendedHeader,
			namespace share.Namespace,
		) (share.NamespacedShares, error) `perm:"read"`
	}
}

func (api *API) SharesAvailable(ctx context.Context, header *header.ExtendedHeader) error {
	return api.Internal.SharesAvailable(ctx, header)
}

func (api *API) GetShare(ctx context.Context, header *header.ExtendedHeader, row, col int) (share.Share, error) {
	return api.Internal.GetShare(ctx, header, row, col)
}

func (api *API) GetEDS(ctx context.Context, header *header.ExtendedHeader) (*rsmt2d.ExtendedDataSquare, error) {
	return api.Internal.GetEDS(ctx, header)
}

func (api *API) GetSharesByNamespace(
	ctx context.Context,
	header *header.ExtendedHeader,
	namespace share.Namespace,
) (share.NamespacedShares, error) {
	return api.Internal.GetSharesByNamespace(ctx, header, namespace)
}

type module struct {
	share.Getter
	share.Availability
}

func (m module) SharesAvailable(ctx context.Context, header *header.ExtendedHeader) error {
	return m.Availability.SharesAvailable(ctx, header)
}