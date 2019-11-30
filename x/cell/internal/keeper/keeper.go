package keeper

import (
	"fmt"
	"github.com/cell-network/cellchain/x/cell/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
)

type Keeper struct {
	stakingKeeper *staking.Keeper
	storeKey      sdk.StoreKey // Unexposed key to access store from sdk.Context
	cdc           *codec.Codec // The wire codec for binary encoding/decoding.
}

func NewKeeper(stakingKeeper *staking.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		stakingKeeper: stakingKeeper,
		storeKey:      storeKey,
		cdc:           cdc,
	}
}

func (k Keeper) GetShare(ctx sdk.Context, valAddr sdk.ValAddress) types.Share {

	fmt.Printf("In getShare: %s", k.storeKey)
	store := ctx.KVStore(k.storeKey)
	if !k.IsValidatorPresent(ctx, valAddr) {
		print("Do not exists")
		return types.NewShare(valAddr)
	}
	print("exists")
	bz := store.Get(valAddr)
	var share types.Share
	k.cdc.MustUnmarshalBinaryBare(bz, &share)
	return share
}

func (k Keeper) SetShare(ctx sdk.Context, valAddr sdk.ValAddress, share types.Share) sdk.Error {
	if share.Owner.Empty() || share.ShareHolders == nil {
		return sdk.ErrInvalidAddress("Validator address does not exist")
	}
	store := ctx.KVStore(k.storeKey)
	store.Set(valAddr, k.cdc.MustMarshalBinaryBare(share))
	return nil
}

func (k Keeper) DeleteShare(ctx sdk.Context, valAddr sdk.ValAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(valAddr)
}

func (k Keeper) AddShareHolder(ctx sdk.Context, valAddr sdk.ValAddress, holder types.ShareHolder) sdk.Error {
	if holder.Address.Empty() || holder.Rate.GTE(sdk.NewDec(1)) || !holder.Rate.IsPositive() {
		return sdk.ErrInternal("Shareholder is invalidate.")
	}

	validator := k.stakingKeeper.Validator(ctx, valAddr)

	if validator == nil {
		return sdk.ErrInternal("Validator does not exists")
	}

	share := k.GetShare(ctx, valAddr)
	total := holder.Rate
	for _, h := range share.ShareHolders {
		if !h.Address.Equals(holder.Address) {
			total = total.Add(h.Rate)
		}
	}
	if total.GTE(sdk.NewDec(1)) {
		return sdk.ErrInternal("Total Sharing Rate can not exceed 100%")
	}
	// replace shareholder if exists
	for i, h := range share.ShareHolders {
		if h.Address.Equals(holder.Address) {
			share.ShareHolders[i] = holder
			k.SetShare(ctx, valAddr, share)
			return nil
		}
	}
	share.ShareHolders = append(share.ShareHolders, holder)
	k.SetShare(ctx, valAddr, share)
	return nil
}

func (k Keeper) DeleteShareHolder(ctx sdk.Context, valAddr sdk.ValAddress, delAddr sdk.AccAddress) sdk.Error {
	if delAddr.Empty() {
		return sdk.ErrInvalidAddress("Delegator Address is invalidate")
	}

	validator := k.stakingKeeper.Validator(ctx, valAddr)

	if validator == nil {
		return sdk.ErrInternal("Validator does not exists")
	}

	share := k.GetShare(ctx, valAddr)

	for i, h := range share.ShareHolders {
		if h.Address.Equals(delAddr) {
			share.ShareHolders = append(share.ShareHolders[:i], share.ShareHolders[i+1:]...)
			k.SetShare(ctx, valAddr, share)
			return nil
		}
	}
	return sdk.ErrInternal("No delegator found.")
}

func (k Keeper) GetValidatorsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

func (k Keeper) IsValidatorPresent(ctx sdk.Context, valAddr sdk.ValAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(valAddr)
}
