package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var Db *gorm.DB = InitDbgGorm()

func InitDbgGorm() *gorm.DB {
	dsn := "admin:tetst@tcp(127.0.0.1:3306)/dsc?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 获取通用数据库对象 sql.DB
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get generic database object")
	}

	// 设置连接池
	sqlDB.SetMaxIdleConns(10)           // 空闲连接池中的最大连接数
	sqlDB.SetMaxOpenConns(100)          // 数据库打开的最大连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接可复用的最大时间
	return db
}
