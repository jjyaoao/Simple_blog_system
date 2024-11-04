package config

import (
	"exchangeapp/global"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() {
	dsn := AppConfig.Database.Dsn

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to initialize database, got error: %v", err)
	}
	sqlDB, err := db.DB()

	sqlDB.SetMaxIdleConns(AppConfig.Database.SetMaxIdleConns)
	sqlDB.SetMaxOpenConns(AppConfig.Database.SetMaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		log.Fatalf("Failed to initialize database, got error: %v", err)
	}

	global.Db = db
}
