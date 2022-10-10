package main

import (
	_ "github.com/joho/godotenv/autoload"

	Db "ip-crawler/db"
	Crawl "ip-crawler/jobs/crawl"
)

func main() {
	Db.Init()
	Db.Migrate()
	Crawl.Run()
}
