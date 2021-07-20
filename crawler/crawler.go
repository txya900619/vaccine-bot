package crawler

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gocolly/colly/v2"
)

type VaccinationInfo struct {
	AstraZeneca string
	Moderna     string
	Total       string
}

type VaccineInventory struct {
	CityName             string
	AstraZenecaRemaining string
	ModernaRemaining     string
}

type VaccineData struct {
	Year                  string
	Month                 string
	Day                   string
	DailyVaccination      VaccinationInfo
	TotalVaccination      VaccinationInfo
	PopulationCoverage    string
	DosagePopulationRatio string
	VaccineInventories    []VaccineInventory
}

func UpdateVaccineData() {
	c := colly.NewCollector()
	c.OnHTML("div.download", func(e *colly.HTMLElement) {

		pdfURL := "https://www.cdc.gov.tw" + e.ChildAttr("p a", "href")

		err := c.Visit(pdfURL)
		if err != nil {
			log.Fatal(err)
			return
		}

	})
	c.OnHTML("a.nav-link.viewer-button", func(e *colly.HTMLElement) {
		downloadURL := "https://www.cdc.gov.tw" + e.Attr("href")

		res, err := http.Get(downloadURL)
		if err != nil {
			log.Fatal(err)
			return
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
			return
		}

		vaccineData, err := parsePDF(body)
		if err != nil {
			log.Fatal(err)
			return
		}

		vaccineDataJSON, err := json.Marshal(vaccineData)
		if err != nil {
			log.Fatal(err)
			return
		}

		err = ioutil.WriteFile("vaccine-data.json", vaccineDataJSON, 0644)
		if err != nil {
			log.Fatal(err)
		}
	})

	err := c.Visit("https://www.cdc.gov.tw/Category/Page/9jFXNbCe-sFK9EImRRi2Og")
	if err != nil {
		log.Fatal(err)
		return
	}
}
