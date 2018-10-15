package color

import (
	"encoding/json"
	"os"
)

type ColorConfig struct {
	BackgroundColor  string `json:"background"`
	WindowTitle      string `json:"windowTitle"`
	WindowBackground string `json:"windowBackground"`
	TextColor        string `json:"text"`
}

var colorConfig ColorConfig

func init() {

	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	jsonParser := json.NewDecoder(file)
	if err := jsonParser.Decode(&colorConfig); err != nil {
		panic(err)
	}

}

func GetColorConfig() ColorConfig {
	return colorConfig
}
