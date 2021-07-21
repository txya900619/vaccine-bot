package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/robfig/cron/v3"
	"github.com/txya900619/vaccine-bot/bot"
	"github.com/txya900619/vaccine-bot/crawler"
)

func main() {
	crawler.UpdateVaccineData()
	crontab := cron.New()
	_, err := crontab.AddFunc("@hourly", func() {
		fmt.Println("crawing")
		crawler.UpdateVaccineData()
	})
	if err != nil {
		log.Fatal(err)
	}

	bot, err := bot.New()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/callback", bot.GetCallbackHandler())

	crontab.Start()

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
