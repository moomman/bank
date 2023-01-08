package main

import (
	"github.com/moonman/mbank/db"
	"github.com/moonman/mbank/token"
	"github.com/moonman/mbank/utils"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/moonman/mbank/global"
	"github.com/moonman/mbank/router"
)

func main() {
	engine := InitRouter()
	Init()
	log.Fatal(engine.Run(global.Config.ServerUrl))
}

func Init() {
	err := utils.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db.Init()
	token.TokenMaker, err = token.NewPasteo_maker(global.Config.Secret)
	if err != nil {
		panic(err)
	}
}

func InitRouter() *gin.Engine {
	engine := gin.Default()
	root := engine.Group("")
	router.Group.Account.Init(root)
	router.Group.User.Init(root)
	return engine
}
