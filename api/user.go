package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/moonman/mbank/db/sqlc"
	"github.com/moonman/mbank/model/request"
	"github.com/moonman/mbank/utils"
)

type user struct{}

func (u *user) CreateUser(c *gin.Context) {
	var req request.CreateUser

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrS(err))
		return
	}
	password, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrS(err))
		return
	}
	_, err = db.Dao.CreateUser(c, &db.CreateUserParams{
		Username:     req.Username,
		HashPassword: password,
		FullName:     req.FullName,
		Email:        req.Email,
		CreateTime:   time.Now(),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrS(err))
		return
	}

	c.JSON(http.StatusOK, nil)
}
