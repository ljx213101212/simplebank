package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func runConcurrentTransaction(n int, store Store, transferParam TransferTxParams, errs chan error, results chan TransferTxResult) {
	// run n concurrent transfer transaction
	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), transferParam)
			errs <- err
			results <- result
		}()
	}
}

func checkTransfer(t *testing.T, store Store, fromAccount Account, toAccount Account, amount int64, transfer Transfer) {
	require.NotEmpty(t, transfer)
	require.Equal(t, fromAccount.ID, transfer.FromAccountID)
	require.Equal(t, toAccount.ID, transfer.ToAccountID)
	require.Equal(t, amount, transfer.Amount)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	_, err := store.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
}

func checkEntry(t *testing.T, store Store, account Account, ammout int64, entry Entry) {

	require.NotEmpty(t, entry)
	require.Equal(t, account.ID, entry.AccountID)
	require.Equal(t, ammout, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	_, err := store.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
}

func checkAccount(t *testing.T, store Store, origin Account, result Account) {
	require.NotEmpty(t, result)
	require.Equal(t, origin.ID, result.ID)
}

func checkBalance(t *testing.T, n int, store Store, originFrom Account, resultFrom Account, originTo Account, resultTo Account, amount int64, existed map[int]bool) {
	fmt.Println(">> tx:", resultFrom.Balance, resultTo.Balance, existed)

	diff1 := originFrom.Balance - resultFrom.Balance
	diff2 := resultTo.Balance - originTo.Balance
	require.Equal(t, diff1, diff2)
	require.True(t, diff1 > 0)
	require.True(t, diff1%amount == 0) // 1 * amount, 2 * amount, 3 * amount, ..., n * amount

	k := int(diff1 / amount)
	require.True(t, k >= 1 && k <= n)
	require.NotContains(t, existed, k)
	existed[k] = true
}

func checkFinalBalance(t *testing.T, n int, store Store, originFrom Account, originTo Account, amount int64) {
	// check the final updated balance
	updatedAccount1, err := store.GetAccount(context.Background(), originFrom.ID)
	require.NoError(t, err)

	updatedAccount2, err := store.GetAccount(context.Background(), originTo.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)

	require.Equal(t, originFrom.Balance-int64(n)*amount, updatedAccount1.Balance)
	require.Equal(t, originTo.Balance+int64(n)*amount, updatedAccount2.Balance)
}

func TestTransferTx(t *testing.T) {

	fmt.Println(">> before NewStore", testDB)
	store := NewStore(testDB)

	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)

	fmt.Println(">> before", fromAccount.Balance, toAccount.Balance)

	// run 5 concurrent transfer transaction
	// transfer 10
	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	// run n concurrent transfer transaction
	runConcurrentTransaction(n, store, TransferTxParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Amount:        amount,
	}, errs, results)
	// check results
	existed := make(map[int]bool)

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		fmt.Println("results:", result)
		require.NotEmpty(t, result)

		//check transfer
		transfer := result.Transfer
		checkTransfer(t, store, fromAccount, toAccount, amount, transfer)

		//check entries
		fromEntry := result.FromEntry
		checkEntry(t, store, fromAccount, -amount, fromEntry)

		toEntry := result.ToEntry
		checkEntry(t, store, toAccount, amount, toEntry)

		// check accounts
		resultFromAccount := result.FromAccount
		checkAccount(t, store, fromAccount, resultFromAccount)

		resultToAccount := result.ToAccount
		checkAccount(t, store, toAccount, resultToAccount)

		// check balances
		checkBalance(t, n, store, fromAccount, resultFromAccount, toAccount, resultToAccount, amount, existed)
	}

	// check the final updated balance
	checkFinalBalance(t, n, store, fromAccount, toAccount, amount)
}
