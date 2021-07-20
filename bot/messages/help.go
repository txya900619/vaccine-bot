package messages

import "github.com/line/line-bot-sdk-go/v7/linebot"

func GetHelpMessage() *linebot.FlexMessage {
	return linebot.NewFlexMessage("help message", &linebot.BubbleContainer{
		Body: &linebot.BoxComponent{
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Size: linebot.FlexTextSizeTypeSm,
					Text: "type /latest to get latest vaccine stat,",
				},
				&linebot.TextComponent{
					Size: linebot.FlexTextSizeTypeSm,
					Text: "type /inventories to get inventory stat,",
				},
				&linebot.TextComponent{
					Size: linebot.FlexTextSizeTypeSm,
					Text: "or click buttom to auto type command",
				},
				&linebot.ButtonComponent{
					Action: linebot.NewMessageAction("/latest", "/latest"),
				},
				&linebot.ButtonComponent{
					Action: linebot.NewMessageAction("/inventories", "/inventories"),
				},
			},
		},
	})
}
