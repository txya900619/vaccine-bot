package messages

import "github.com/line/line-bot-sdk-go/v7/linebot"

func GetDefaultMessage() *linebot.TextMessage {
	return linebot.NewTextMessage(
		"Welcome to vaccine bot,\nthis bot will tell you Covid-19 vaccine statistical data,\ntype /help to get all avaliable command to use.",
	)

}
