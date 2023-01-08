package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/moonman/mbank/db/sqlc"
	"github.com/moonman/mbank/model/request"
	"github.com/moonman/mbank/utils"
)

type account struct{}

func (a *account) Transfer(c *gin.Context) {
	var req request.Transfer

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	err := db.Dao.TransferTo(c, &db.TransferToParams{
		From:   req.From,
		To:     req.To,
		Amount: req.Amount,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrS(err))
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (a *account) AddAccount(c *gin.Context) {
	var req request.AddAccount

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrS(err))
	}

	_, err := db.Dao.CreateAccount(c, &db.CreateAccountParams{
		Owner:     req.Owner,
		Balance:   req.Balance,
		Currency:  req.Currency,
		CreatedAt: time.Now(),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrS(err))
		return
	}

	c.JSON(http.StatusOK, nil)
}
