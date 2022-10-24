package main

import (
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/takeuchi-shogo/line-bot-sample/config"
)

func main() {

	c := config.NewConfig()
	// LINE Botクライアント生成する
	// BOT にはチャネルシークレットとチャネルトークンを環境変数から読み込み引数に渡す
	bot, err := linebot.New(
		c.ChannelSecret,
		c.ChannelToken,
	)

	if err != nil {
		log.Fatal(err)
	}

	message := linebot.NewTextMessage("Hello world")
	// テキストメッセージを友達登録しているユーザー全員に配信する
	if _, err = bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}
