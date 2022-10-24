package main

import (
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/takeuchi-shogo/line-bot-sample/config"
)

func main() {

	c := config.NewConfig()

	bot, err := linebot.New(
		c.ChannelSecret,
		c.ChannelToken,
	)

	if err != nil {
		log.Fatal(err)
	}

	message := linebot.NewTextMessage("Hello world")

	if _, err = bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}
