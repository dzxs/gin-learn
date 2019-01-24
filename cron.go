package main

import (
	"github.com/dzxs/blog/models"
	"github.com/robfig/cron"
	"log"
	"time"
)

func main() {
	log.Println("Strating...")
	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllArticle()
	})
	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
