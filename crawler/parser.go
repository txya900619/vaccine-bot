package crawler

import (
	"bytes"
	"fmt"
	"io"
	"regexp"

	"github.com/ledongthuc/pdf"
)

var cityNames = [23]string{
	"臺北市",
	"新北市",
	"桃園市",
	"臺中市",
	"臺南市",
	"高雄市",
	"新竹縣",
	"彰化縣",
	"雲林縣",
	"屏東縣",
	"基隆市",
	"宜蘭縣",
	"新竹市",
	"苗栗縣",
	"嘉義市",
	"嘉義縣",
	"花蓮縣",
	"臺東縣",
	"南投縣",
	"澎湖縣",
	"金門縣",
	"連江縣",
	"總計",
}

func readPDF(raw []byte) (string, error) {
	pdfReader, err := pdf.NewReader(bytes.NewReader(raw), int64(len(raw)))
	if err != nil {
		return "", err
	}

	contentReader, err := pdfReader.GetPlainText()
	if err != nil {
		return "", err
	}

	b, err := io.ReadAll(contentReader)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func parsePDF(raw []byte) (*VaccineData, error) {

	content, err := readPDF(raw)
	if err != nil {
		return nil, err
	}

	regex, err := regexp.Compile(`\d+([,|.]?\d*)*`)
	if err != nil {
		return nil, err
	}

	numbers := regex.FindAllString(content, -1)
	fmt.Println(numbers)

	vaccineInventories := make([]VaccineInventory, 0)

	for i, cityName := range cityNames {
		vaccineInventories = append(vaccineInventories, VaccineInventory{
			CityName:             cityName,
			AstraZenecaRemaining: numbers[60+11*i],
			ModernaRemaining:     numbers[65+11*i],
		})
	}

	return &VaccineData{
		Year:  numbers[1],
		Month: numbers[2],
		Day:   numbers[3],
		DailyVaccination: VaccinationInfo{
			AstraZeneca: numbers[10],
			Moderna:     numbers[15],
			Total:       numbers[9],
		},
		TotalVaccination: VaccinationInfo{
			AstraZeneca: numbers[21],
			Moderna:     numbers[22],
			Total:       numbers[20],
		},
		PopulationCoverage:    numbers[24],
		DosagePopulationRatio: numbers[25],
		VaccineInventories:    vaccineInventories,
	}, nil
}
