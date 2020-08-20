package main

import (
	"config"
	"fmt"
	"github.com/gin-gonic/gin"
	"model"
	"os"
	"router"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	err := config.Init("./config.json")
	if err != nil {
		fmt.Println("config.Init err: ", err)
		os.Exit(500)
	}
	err = model.Init(config.Config.MysqlSetting[config.SrvName].MysqlConn, config.Config.MysqlSetting[config.SrvName].MysqlConnectPoolSize)
	if err != nil {
		fmt.Println("model.Init err: ", err)
		os.Exit(500)
	}
	fmt.Println(config.SrvName, " starting...")
	engine := gin.New()
	router.GroupRouter(engine)
	engine.Run(config.Config.SrvSetting[config.SrvName].SrvListen)
}
