package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Basebuild17/cosmos-blockchain/x/registry/types"
)

type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey storetypes.StoreKey
	memKey   storetypes.StoreKey
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetData sets a key-value pair in the registry
func (k Keeper) SetData(ctx sdk.Context, key string, value string) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(key), []byte(value))
}

// GetData retrieves a value by key from the registry
func (k Keeper) GetData(ctx sdk.Context, key string) (string, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(key))
	if bz == nil {
		return "", false
	}
	return string(bz), true
}

// DeleteData deletes a key-value pair from the registry
func (k Keeper) DeleteData(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(key))
}

// GetAllData returns all data in the registry
func (k Keeper) GetAllData(ctx sdk.Context) []types.Data {
	store := ctx.KVStore(k.storeKey)
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	var data []types.Data
	for ; iterator.Valid(); iterator.Next() {
		data = append(data, types.Data{
			Key:   string(iterator.Key()),
			Value: string(iterator.Value()),
		})
	}
	return data
}
