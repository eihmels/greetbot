package fileloader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	Bot struct {
		Name  string `json:"name"`
		Token string `json:"token"`
	} `json:"bot"`
	Channel      string `json:"channel"`
	IgnoredUsers []struct {
		Name string `json:"name"`
	} `json:"ignored-users"`
	GreetingsFile  string `json:"greeting-file"`
	SalutionFile   string `json:"Salution-file"`
	ComplimentFile string `json:"compliment-file"`
}

func LoadConfigFromJsonFile() (config Config) {
	fmt.Println("load config from ../config/conf.json")

	jsonFile, err := os.Open("config/conf.json")

	if err != nil {
		panic(err)
	}

	byteValue, _ := io.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &config)

	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	return config
}
