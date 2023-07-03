package client

import (
	"github.com/eihmels/greetbot/fileloader"
	"github.com/eihmels/greetbot/tools"
	"github.com/gempir/go-twitch-irc/v4"
	"log"
	"math/rand"
	"time"
)

func GetClient(config fileloader.Config) *twitch.Client {
	client := twitch.NewClient(config.Bot.Name, config.Bot.Token)

	var joinedUsers []string

	timestamp := time.Now().Unix()

	for _, IgnoredUsers := range config.IgnoredUsers {
		println("add", IgnoredUsers.Name, "to IgnoredList")
		joinedUsers = append(joinedUsers, IgnoredUsers.Name)
	}

	client.OnPingMessage(func(message twitch.PingMessage) {
		result := time.Now().Unix() - timestamp
		log.Println(result)

		if result > config.AskingPeriod {
			botMessage := getARandomLine(config.GreetingsFile)

			log.Println("say: " + botMessage)
			client.Say(config.Channel, botMessage)

			timestamp = time.Now().Unix()
		}
	})

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		timestamp = time.Now().Unix()

		if tools.StringInSlice(message.User.Name, joinedUsers) == false {
			botMessage := createMessage(message.User.Name, config)

			log.Println("say: " + botMessage)
			client.Say(config.Channel, botMessage)

			log.Println(message.User.Name, " added to IgnoredList")
			joinedUsers = append(joinedUsers, message.User.Name)
		}
	})

	client.Join(config.Channel)

	return client
}

func createMessage(user string, config fileloader.Config) string {
	return getARandomLine(config.SalutationFile) + " @" + user + " " + getARandomLine(config.ComplimentFile) + " " + getARandomLine(config.GreetingsFile)
}

func getARandomLine(path string) string {
	slice, err := fileloader.LoadGreetings(path)

	if err != nil {
		panic(err)
	}

	return slice[GetRandomNumber(len(slice))]
}

func GetRandomNumber(n int) (number int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	return r1.Intn(n)
}
