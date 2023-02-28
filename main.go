package main

import (
	"github.com/eihmels/greetbot/fileloader"
	"log"
	"math/rand"
	"time"

	twitch "github.com/gempir/go-twitch-irc/v4"
)

func main() {
	config := fileloader.LoadConfigFromJsonFile()

	client := twitch.NewClient(config.Bot.Name, config.Bot.Token)

	var joinedUsers []string

	for _, IgnoredUsers := range config.IgnoredUsers {
		println("add ", IgnoredUsers.Name, "to IgnoredList")
		//joinedUsers = append(joinedUsers, IgnoredUsers.Name)
	}

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		if stringInSlice(message.User.Name, joinedUsers) == false {
			botMessage := getARandomLine(config.SalutionFile) + " @" + message.User.Name + " " + getARandomLine(config.ComplimentFile) + " " + getARandomLine(config.GreetingsFile)

			log.Println("say: " + botMessage)
			client.Say(config.Channel, botMessage)

			log.Println(message.User.Name, " added to IgnoredList")
			joinedUsers = append(joinedUsers, message.User.Name)
		}
	})

	client.Join(config.Channel)

	err := client.Connect()

	if err != nil {
		panic(err)
	}
}

func getARandomLine(path string) string {
	slice, err := fileloader.LoadGreetings(path)

	if err != nil {
		panic(err)
	}

	return slice[GetRandomNumber(len(slice))]
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {

			return true
		}
	}
	return false
}

func GetRandomNumber(n int) (number int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	return r1.Intn(n)
}
