package initialize

import (
	"go.uber.org/zap"
	"shortener/lib/db"
	"shortener/models"
)

func InitDB() {
	db.SetupDBConn()
	db := db.GetDB()
	err := db.AutoMigrate(&models.UrlModel{})
	if err != nil {
		zap.S().Error("init table error.")
	}
}
