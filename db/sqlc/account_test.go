package db

import (
	"context"
	"testing"
	"time"

	"github.com/moonman/mbank/utils"
	"github.com/stretchr/testify/require"
)

//实现一个实例操作：从account1转10￥到account2：
//创建tranfer记录
//创建entry -10 +10
//对账户余额进行操作
func createRandomAccount() *Account {
	params := Account{
		Owner:     utils.RandomOwner(),
		Balance:   1000,
		Currency:  "USD",
		CreatedAt: time.Now(),
	}
	res, err := testQueries.CreateAccount(context.Background(), &CreateAccountParams{
		Owner:     params.Owner,
		Balance:   params.Balance,
		Currency:  params.Currency,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return nil
	}
	params.ID, _ = res.LastInsertId()
	return &params
}

func TestAccount(t *testing.T) {
	//err := testQueries.CreateAccount(context.Background(), &CreateAccountParams{
	//	Owner:     "ml",
	//	Balance:   1000,
	//	Currency:  "USD",
	//	CreatedAt: time.Now(),
	//})
	//require.NoError(t, err)
	//
	//account, err := testQueries.GetAccountById(context.Background(), 1)
	//require.NoError(t, err)
	//log.Println(account)
}

func TestAccountTx(t *testing.T) {

	a1 := createRandomAccount()
	a2 := createRandomAccount()

	n := 5 //循环次数
	amount := 100
	errCh := make(chan error, n)
	for i := 0; i < n; i++ {
		go func() {
			err := testQueries.TransferTo(context.Background(), &TransferToParams{
				From:   a1.ID,
				To:     a2.ID,
				Amount: int64(amount),
			})
			errCh <- err
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errCh
		require.NoError(t, err)
	}

	account1, err := testQueries.GetAccountById(context.Background(), a1.ID)
	require.NoError(t, err)
	account2, err := testQueries.GetAccountById(context.Background(), a2.ID)
	require.NoError(t, err)

	require.Equal(t, a1.Balance-int64(n*amount), account1.Balance)
	require.Equal(t, int64(n*amount)+a2.Balance, account2.Balance)
}
