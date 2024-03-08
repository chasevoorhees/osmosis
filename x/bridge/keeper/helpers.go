package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"golang.org/x/exp/slices"
)

// Difference returns the slice of elements that are elements of a but not elements of b.
// TODO: Placed here temporarily. Delete after releasing the new osmoutils version.
func Difference[T comparable](a, b []T) []T {
	mb := make(map[T]struct{}, len(a))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	diff := make([]T, 0)
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

// validateSigners ensures that the actualSigners set
// is a non-empty subset of the validSigners set
func (k Keeper) validateSender(ctx sdk.Context, sender string) bool {
	return slices.Contains(k.GetParams(ctx).Signers, sender)
}
