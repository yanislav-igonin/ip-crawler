// interval job to crawl the web for new content
package jobs

import (
	"ip-crawler/db"
	"ip-crawler/http"
	"ip-crawler/ip"
	"ip-crawler/models"
	"time"

	"github.com/robfig/cron"
)

func crawl() {
	address := ip.GetRandom()
	var inDb models.Ip
	db.Client.First(&inDb, "address = ?", address)
	if inDb.Address != "" {
		return
	}
	status, statusCode := http.Request(address)
	db.Client.Create(&models.Ip{
		Address:    address,
		Status:     status,
		CheckedAt:  time.Now(),
		StatusCode: statusCode,
	})
}

func Run() {
	c := cron.New()
	c.AddFunc("@every 1s", crawl)
	c.Run()
}
