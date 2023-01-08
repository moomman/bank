package router

import (
	"github.com/gin-gonic/gin"
	"github.com/moonman/mbank/api"
	"github.com/moonman/mbank/middleware"
)

type user struct {
}

func (user) Init(g *gin.RouterGroup) {
	group := g.Group("user")
	{
		manger := group.Group("").Use(middleware.Auth())
		group.POST("login", api.Group.User.Login)
		manger.POST("create", api.Group.User.CreateUser)

		//group.POST("add", api.Group.Account.AddAccount)
	}
}
