package db

import (
	"context"
	"testing"
	"time"

	"github.com/JuanHeredia3/simple-bank/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func CreateRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: pgtype.Int8{Int64: account.ID, Valid: true},
		Amount:    util.RandomMoney(),
	}
	entry, err := testStore.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID.Int64, entry.AccountID.Int64)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account := CreateRandomAccount(t)
	CreateRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := CreateRandomAccount(t)
	entry1 := CreateRandomEntry(t, account)
	entry2, err := testStore.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestDeleteEntry(t *testing.T) {
	account := CreateRandomAccount(t)
	entry1 := CreateRandomEntry(t, account)
	err := testStore.DeleteEntry(context.Background(), entry1.ID)
	require.NoError(t, err)

	entry2, err := testStore.GetEntry(context.Background(), entry1.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, entry2)
}

func TestUpdateEntry(t *testing.T) {
	account := CreateRandomAccount(t)
	entry1 := CreateRandomEntry(t, account)
	arg := UpdateEntryParams{
		ID:     entry1.ID,
		Amount: util.RandomMoney(),
	}

	entry2, err := testStore.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, arg.Amount, entry2.Amount)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	account := CreateRandomAccount(t)

	for range 10 {
		CreateRandomEntry(t, account)
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testStore.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, account.ID, entry.AccountID.Int64)
	}
}
