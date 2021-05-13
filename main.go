package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"shortener/global"
	"shortener/initialize"
	"syscall"
	"time"
)

func main() {
	initialize.InitLogger()
	initialize.InitConfig()
	if err := initialize.InitTrans(global.ServerConfig.UserConfig.Local); err != nil {
		zap.S().Infof("初始化翻译器错误: %s", global.ServerConfig.UserConfig.Local)
		return
	}
	initialize.InitDB()
	router := initialize.Router()
	addr, port := global.ServerConfig.UserConfig.Host, global.ServerConfig.UserConfig.Port
	zap.S().Infof("启动地址：%s，端口：%d", addr, port)
	router.Run(fmt.Sprintf("%s:%d", addr, port))

	server := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", addr, port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		// service connections
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.S().Fatalf("listen: %s\n", err)
		}
	}()
	gracefulExitWeb(server)
}

func gracefulExitWeb(server *http.Server) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	sig := <-ch

	zap.S().Info("got a signal", sig)
	now := time.Now()
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(cxt)
	if err != nil {
		zap.S().Fatal("err", err)
	}

	// 看看实际退出所耗费的时间
	zap.S().Info("------exited--------", time.Since(now))
}
