package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/moonman/mbank/db"
	"github.com/moonman/mbank/global"
	"github.com/moonman/mbank/router"
	"github.com/moonman/mbank/utils"
)

func main() {
	engine := InitRouter()
	err := utils.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db.Init()
	log.Fatal(engine.Run(global.Config.ServerUrl))
}

func InitRouter() *gin.Engine {
	engine := gin.Default()
	root := engine.Group("")
	router.Group.Account.Init(root)

	return engine
}
