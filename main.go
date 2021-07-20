package main

import (
	"log"
	"net/http"
	"os"

	"github.com/robfig/cron/v3"
	"github.com/txya900619/vaccine-bot/bot"
	"github.com/txya900619/vaccine-bot/crawler"
)

func main() {
	crawler.UpdateVaccineData()
	crontab := cron.New(cron.WithLogger(cron.DefaultLogger))
	crontab.AddFunc("@hourly", crawler.UpdateVaccineData)

	bot, err := bot.New()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/callback", bot.GetCallbackHandler())

	crontab.Start()
	defer crontab.Stop()

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
