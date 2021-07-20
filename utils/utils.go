package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/txya900619/vaccine-bot/crawler"
)

func ReadVaccineData() (*crawler.VaccineData, error) {
	vaccineDataJSON, err := ioutil.ReadFile("vaccine-data.json")
	if err != nil {
		return nil, err
	}

	vaccineData := &crawler.VaccineData{}
	err = json.Unmarshal(vaccineDataJSON, vaccineData)
	if err != nil {
		return nil, err
	}

	return vaccineData, nil
}
