package messages

import (
	"log"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/txya900619/vaccine-bot/utils"
)

func GetInventoriesMessage() *linebot.FlexMessage {
	vaccineData, err := utils.ReadVaccineData()
	if err != nil {
		log.Fatal(err)
	}

	vaccineInventoriesBoxes := make([]linebot.FlexComponent, 0)

	vaccineInventoriesBoxes = append(vaccineInventoriesBoxes, horizontalThreeTextBox("縣市別", "AZ 剩餘劑量", "莫德納剩餘劑量"))

	for _, vaccineInventory := range vaccineData.VaccineInventories {
		vaccineInventoriesBoxes = append(vaccineInventoriesBoxes, horizontalThreeTextBox(vaccineInventory.CityName, vaccineInventory.AstraZenecaRemaining, vaccineInventory.ModernaRemaining))
	}

	return linebot.NewFlexMessage("vaccine inventories info", &linebot.BubbleContainer{
		Size: linebot.FlexBubbleSizeTypeGiga,
		Body: &linebot.BoxComponent{
			Layout:   linebot.FlexBoxLayoutTypeVertical,
			Contents: vaccineInventoriesBoxes,
		},
	})
}
