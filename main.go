package main

import (
	_ "github.com/joho/godotenv/autoload"

	"ip-crawler/config"
	"ip-crawler/db"
	crawl "ip-crawler/jobs/crawl"
)

func main() {
	config.Init()
	db.Init()
	db.Migrate()
	crawl.Run()
}
