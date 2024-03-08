package keeper

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/osmosis/v23/x/bridge/types"
)

type UpdateParamsResult struct {
	signersToCreate []string
	signersToDelete []string
	assetsToCreate  []types.AssetWithStatus
	assetsToDelete  []types.AssetWithStatus
}

// UpdateParams properly updates params of the module.
func (k Keeper) UpdateParams(ctx sdk.Context, newParams types.Params) (UpdateParamsResult, error) {
	var (
		oldParams = k.GetParams(ctx)

		signersToCreate = Difference(newParams.Signers, oldParams.Signers)
		signersToDelete = Difference(oldParams.Signers, newParams.Signers)
		assetsToCreate  = Difference(newParams.Assets, oldParams.Assets)
		assetsToDelete  = Difference(oldParams.Assets, newParams.Assets)
	)

	// create denoms for all new assets
	err := k.createAssets(ctx, assetsToCreate)
	if err != nil {
		return UpdateParamsResult{}, err
	}

	// disable deleted assets
	for _, asset := range assetsToDelete {
		_, err = k.ChangeAssetStatus(ctx, asset.Asset, types.AssetStatus_ASSET_STATUS_BLOCKED_BOTH)
		if err != nil {
			return UpdateParamsResult{},
				errorsmod.Wrapf(types.ErrCantChangeAssetStatus, "Can't disable asset %v: %s", asset.Asset, err)
		}
	}

	// don't need to specifically update the signers, just save them

	k.SetParams(ctx, newParams)

	return UpdateParamsResult{
		signersToCreate: signersToCreate,
		signersToDelete: signersToDelete,
		assetsToCreate:  assetsToCreate,
		assetsToDelete:  assetsToDelete,
	}, nil
}

// SetParams sets the total set of params.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// GetParams returns the total set params.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}
