package db

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"shortener/global"
	"time"
)

// db连接
var db *gorm.DB

func NewConnection() *gorm.DB {
	//db = newConnection()
	var dbURI string
	var dialector gorm.Dialector
	if global.ServerConfig.DatabaseConfig.Type == "mysql" {
		dbURI = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
			global.ServerConfig.DatabaseConfig.User,
			global.ServerConfig.DatabaseConfig.Password,
			global.ServerConfig.DatabaseConfig.Host,
			global.ServerConfig.DatabaseConfig.Port,
			global.ServerConfig.DatabaseConfig.Name)
		dialector = mysql.New(mysql.Config{
			DSN:                       dbURI, // data source name
			DefaultStringSize:         256,   // default size for string fields
			DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
			DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
			DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
			SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
		})
		zap.S().Info("启用Mysql数据库")

	} else if global.ServerConfig.DatabaseConfig.Type == "postgres" {
		dbURI = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
			global.ServerConfig.DatabaseConfig.Host,
			global.ServerConfig.DatabaseConfig.Port,
			global.ServerConfig.DatabaseConfig.User,
			global.ServerConfig.DatabaseConfig.Name,
			global.ServerConfig.DatabaseConfig.Password)
		dialector = postgres.New(postgres.Config{
			DSN:                  "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		})
	} else { // sqlite3
		file, _ := os.Getwd()
		dbPath := fmt.Sprintf("%s/%s", file, "test.db")
		zap.S().Infof("启用sqlite数据库：%s", dbPath)
		dialector = sqlite.Open(dbPath)
	}
	conn, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		zap.S().Error(err.Error())
	}

	return conn
}

func SetupDBConn() {
	db = NewConnection()

	sqlDB, err := db.DB()
	if err != nil {
		zap.S().Error("connect db server failed.")
	}
	sqlDB.SetMaxIdleConns(global.ServerConfig.DatabaseConfig.MinConn) // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxOpenConns(global.ServerConfig.DatabaseConfig.MaxConn) // SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetConnMaxLifetime(time.Second * 600)           // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
}

func GetDB() *gorm.DB {
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Errorf("connect db server failed.")
		NewConnection()
	}
	if err := sqlDB.Ping(); err != nil {
		sqlDB.Close()
		NewConnection()
	}
	return db
}
