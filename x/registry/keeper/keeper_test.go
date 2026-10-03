package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cometbft/cometbft-db"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cometbft/cometbft/libs/log"

	registrytypes "github.com/Basebuild17/cosmos-blockchain/x/registry/types"
)

func TestKeeperDeleteData(t *testing.T) {
	db := dbm.NewMemDB()
	storeKey := storetypes.NewKVStoreKey(registrytypes.StoreKey)
	memKey := storetypes.NewMemoryStoreKey(registrytypes.MemStoreKey)
	multiStore := store.NewCommitMultiStore(db, log.NewNopLogger())
	require.NoError(t, multiStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db))
	require.NoError(t, multiStore.LoadLatestVersion())

	ctx := sdk.NewContext(multiStore, types.Header{}, false, log.NewNopLogger())
	keeper := NewKeeper(
		codec.NewProtoCodec(codectypes.NewInterfaceRegistry()),
		storeKey,
		memKey,
	)

	keeper.SetData(ctx, "remove-me", "value")
	keeper.SetData(ctx, "keep-me", "value")

	keeper.DeleteData(ctx, "remove-me")

	_, found := keeper.GetData(ctx, "remove-me")
	require.False(t, found)

	value, found := keeper.GetData(ctx, "keep-me")
	require.True(t, found)
	require.Equal(t, "value", value)
}
