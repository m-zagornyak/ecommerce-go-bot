package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/m-zagornyak/ecommerce-go-bot/pkg/telegram"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5347241093:AAHhBE-pyCpUDQitwmbZG0BNnLsnbJg5u2E")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegramBot := telegram.Newbot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
