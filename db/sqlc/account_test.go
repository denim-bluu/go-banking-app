package db

import (
	"context"
	"testing"

	"github.com/denim-bluu/simplebank/util"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{Owner: util.RandOwner(), Balance: util.RandomMoney(), Currency: util.RnadomCurrency()}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)

	require.Equal(t, account.Owner, arg.Owner)
	require.Equal(t, account.Balance, arg.Balance)
	require.Equal(t, account.Currency, arg.Currency)

	require.NotEmpty(t, account.ID)
	require.NotEmpty(t, account.CreatedAt)
	return account
}
func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2, err := testQueries.GetAccount(context.Background(), acc1.ID)

	require.NoError(t, err)

	require.Equal(t, acc1, acc2)
}
func TestUpdateAccount(t *testing.T) {
	acc1 := createRandomAccount(t)

	arg := UpdateAccountParams{ID: acc1.ID, Balance: util.RandomMoney()}
	acc2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)

	acc3, err := testQueries.GetAccount(context.Background(), acc1.ID)
	require.NoError(t, err)

	require.NotEqual(t, acc1, acc2)
	require.Equal(t, acc2, acc3)
}

func TestDeleteAccount(t *testing.T) {
	acc1 := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), acc1.ID)
	require.NoError(t, err)

	acc2, err := testQueries.GetAccount(context.Background(), acc1.ID)

	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, acc2)
}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{5, 5}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, accounts, 5)
	for _, acc := range accounts {
		require.NotEmpty(t, acc)
	}
}
