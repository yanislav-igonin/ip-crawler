package main

import (
	Db "ip-crawler/db"
	Crawl "ip-crawler/jobs/crawl"
)

func main() {
	Db.Init()
	Db.Migrate()
	Crawl.Run()
}
