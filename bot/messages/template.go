package messages

import "github.com/line/line-bot-sdk-go/v7/linebot"

func centalBoldText(prefix string, central string, suffix string) *linebot.TextComponent {
	return &linebot.TextComponent{
		Contents: []*linebot.SpanComponent{
			{
				Text: prefix,
			},
			{
				Text:   central,
				Weight: linebot.FlexTextWeightTypeBold,
			},
			{
				Text: suffix,
			},
		},
	}
}

func horizontalThreeTextBox(first string, second string, third string) *linebot.BoxComponent {
	return &linebot.BoxComponent{
		Layout: linebot.FlexBoxLayoutTypeHorizontal,
		Contents: []linebot.FlexComponent{
			&linebot.TextComponent{
				Text:  first,
				Align: linebot.FlexComponentAlignTypeCenter,
			},
			&linebot.TextComponent{
				Text:  second,
				Align: linebot.FlexComponentAlignTypeCenter,
			},
			&linebot.TextComponent{
				Text:  third,
				Align: linebot.FlexComponentAlignTypeCenter,
			},
		},
	}
}
