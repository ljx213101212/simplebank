package db_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/util"

	_ "simplebank/db"
)

func TestCreateAccount(t *testing.T) {

	arg := CreateAccountParams{
		Owner:    "tom",
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	user, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

}
