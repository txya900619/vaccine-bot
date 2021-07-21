package messages

import (
	"fmt"
	"log"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/txya900619/vaccine-bot/utils"
)

func GetLatestMessage() *linebot.FlexMessage {
	vaccineData, err := utils.ReadVaccineData()
	if err != nil {
		log.Fatal(err)
	}

	return linebot.NewFlexMessage("latest covid-19 vaccine info", &linebot.BubbleContainer{
		Size: linebot.FlexBubbleSizeTypeMega,
		Body: &linebot.BoxComponent{
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Text: fmt.Sprintf("%s/%s/%s COVID-19 疫苗數據", vaccineData.Year, vaccineData.Month, vaccineData.Day),
					Size: linebot.FlexTextSizeTypeLg,
				},
				centalBoldText("當日共有 ", vaccineData.DailyVaccination.Total, " 接種人次，"),
				centalBoldText("其中 AZ 佔 ", vaccineData.DailyVaccination.AstraZeneca, " 人次，"),
				centalBoldText("莫德納佔 ", vaccineData.DailyVaccination.Moderna, " 人次，"),
				&linebot.SeparatorComponent{
					Margin: linebot.FlexComponentMarginTypeLg,
				},
				centalBoldText("累計接種 ", vaccineData.TotalVaccination.Total, " 人次，"),
				centalBoldText("其中 AZ 佔 ", vaccineData.TotalVaccination.AstraZeneca, " 人次，"),
				centalBoldText("莫德納佔 ", vaccineData.TotalVaccination.Moderna, " 人次，"),
				centalBoldText("接種人口覆蓋率 ", vaccineData.PopulationCoverage+"%", ","),
				centalBoldText("劑次人口比 ", vaccineData.DosagePopulationRatio, " (劑 /每百人 )。"),
				&linebot.ButtonComponent{
					Action: &linebot.MessageAction{
						Label: "點此觀看各縣市剩餘劑量",
						Text:  "/inventories",
					},
				},
			},
		},
	})
}
