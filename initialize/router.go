package initialize

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	routers "shortener/router"
)

func Router() *gin.Engine {
	zap.S().Info("开始启动服务...")
	router := gin.Default()
	//router.LoadHTMLGlob("templates/*")
	routers.InitApi(router)
	return router
}
