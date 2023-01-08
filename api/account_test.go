package api

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	mockdb "github.com/moonman/mbank/db/mock"
	db "github.com/moonman/mbank/db/sqlc"
	"github.com/moonman/mbank/utils"
)

func TestCreateAccount(t *testing.T) {
	account := randomAccount()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)
	store.EXPECT().GetAccountById(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(account, nil)

}

func randomAccount() *db.Account {
	return &db.Account{
		ID:        1,
		Owner:     utils.RandomOwner(),
		Balance:   1000,
		Currency:  "USD",
		CreatedAt: time.Now(),
	}
}
