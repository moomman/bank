package router

import (
	"github.com/gin-gonic/gin"
	"github.com/moonman/mbank/api"
)

type account struct {
}

func (account) Init(g *gin.RouterGroup) {
	group := g.Group("account")
	{
		group.POST("transfer", api.Group.Account.Transfer)
	}
}
