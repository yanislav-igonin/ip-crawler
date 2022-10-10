package main

import (
	"fmt"
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

	fmt.Println("Generating IPs...")
	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			for k := 0; k < 256; k++ {
				var ips []models.Ip

				for l := 0; l < 256; l++ {
					address := ip.Generate(i, j, k, l)
					ips = append(ips, models.Ip{Address: address})
				}
				db.Create(&ips)
			}
			fmt.Printf("Range %d.%d.x.x saved!\n", i, j)
		}
		fmt.Printf("Range %d.x.x.x saved!\n", i)
	}
}
