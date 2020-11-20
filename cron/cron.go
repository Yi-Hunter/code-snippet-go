package main

import (
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()
	id, err := c.AddFunc("*/1 * * * *", func() {
		log.Println(time.Now())
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(id)
	c.Start()
	select {}
}
