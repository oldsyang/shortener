package main

import (
	"fmt"
	"go.uber.org/zap"
	"shortener/global"
	"shortener/initialize"
)

func main() {
	initialize.InitLogger()
	initialize.InitConfig()
	if err := initialize.InitTrans(global.ServerConfig.UserConfig.Local); err != nil {
		fmt.Println("初始化翻译器错误")
		return
	}
	initialize.InitDB()
	router := initialize.Router()
	addr, port := global.ServerConfig.UserConfig.Host, global.ServerConfig.UserConfig.Port
	zap.S().Infof("启动地址：%s，端口：%d", addr, port)
	router.Run(fmt.Sprintf("%s:%d", addr, port))
}
