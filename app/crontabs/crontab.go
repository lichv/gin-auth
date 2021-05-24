package crontabs

import (
	"github.com/robfig/cron"
	"log"
)


func Setup() *cron.Cron {
	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
	})

	return c
}
