package main

import (
	"ip-crawler/http"
	"ip-crawler/ip"
	"ip-crawler/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("local.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Ip{})

	address := ip.GetRandom()
	var inDb models.Ip
	db.First(&inDb, "address = ?", address)
	if inDb.Address != "" {
		return
	}
	status := http.Request(address)
	db.Create(&models.Ip{Address: address, Status: status})
}
