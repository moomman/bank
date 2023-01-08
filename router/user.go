package router

import (
	"github.com/gin-gonic/gin"
	"github.com/moonman/mbank/api"
)

type user struct {
}

func (user) Init(g *gin.RouterGroup) {
	group := g.Group("user")
	{
		group.POST("create", api.Group.User.CreateUser)
		//group.POST("add", api.Group.Account.AddAccount)
	}
}
