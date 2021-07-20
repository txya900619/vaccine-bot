package bot

import (
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/txya900619/vaccine-bot/bot/messages"
)

type VaccineBot struct {
	*linebot.Client
}

func New() (*VaccineBot, error) {
	bot, err := linebot.New(os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_TOKEN"))
	if err != nil {
		return nil, err
	}

	return &VaccineBot{bot}, nil
}

func (bot *VaccineBot) GetCallbackHandler() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				res.WriteHeader(400)
			} else {
				res.WriteHeader(500)
			}
			return
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					var replyMessage linebot.SendingMessage = messages.GetDefaultMessage()
					switch message.Text {
					case "/help":
						replyMessage = messages.GetHelpMessage()
					case "/latest":
						replyMessage = messages.GetLatestMessage()
					case "/inventories":
						replyMessage = messages.GetInventoriesMessage()
					}

					reply := bot.ReplyMessage(event.ReplyToken, replyMessage)
					_, err = reply.Do()
					if err != nil {
						log.Fatal(err)
					}
				}

			}
		}

	}
}
