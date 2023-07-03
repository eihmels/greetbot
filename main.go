package main

import (
	"github.com/eihmels/greetbot/client"
	"github.com/eihmels/greetbot/fileloader"
)

func main() {
	config := LoadConfig()

	twitchClient := client.GetClient(config)

	err := twitchClient.Connect()

	if err != nil {
		panic(err)
	}
}

func LoadConfig() fileloader.Config {
	return fileloader.LoadConfigFromJsonFile()
}
