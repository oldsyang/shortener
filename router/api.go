package router

import (
	"github.com/gin-gonic/gin"
	api "shortener/core"
)

func InitApi(router *gin.Engine) {
	routerGroup := router.Group("/api/")
	{
		routerGroup.POST("/encode", api.Encode)
		routerGroup.GET("/info", api.Info)
	}
	router.GET("/3/:code", api.Redirect)
}
