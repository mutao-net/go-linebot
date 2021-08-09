package main

import (
	"log"
	"os"
	"strconv"

	"github.com/mutao-net/go-linebot/wine"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Print(err)
	}

	results := wine.GetRecommendResults()

	for _, item := range results.Items {
		var messagese string
		messagese = messagese + item.Item.ItemName + "\n"
		messagese = messagese + "¥" + strconv.Itoa(item.Item.ItemPrice) + "\n"
		messagese = messagese + item.Item.ItemURL
		message := linebot.NewTextMessage(messagese)
		if _, err := bot.BroadcastMessage(message).Do(); err != nil {
			log.Print(err)
		}
	}
}
