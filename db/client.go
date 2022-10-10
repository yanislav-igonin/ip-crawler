package db

import (
	"ip-crawler/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Client *gorm.DB

func Init() {
	db, err := gorm.Open(sqlite.Open("local.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Client = db
}

func Migrate() {
	err := Client.AutoMigrate(&models.Ip{})
	if err != nil {
		panic("failed to migrate database")
	}
}
