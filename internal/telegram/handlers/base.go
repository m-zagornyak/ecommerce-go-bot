package handlers

import (
	tele "gopkg.in/telebot.v3"
)

func RegisterBaseHandlers(bot *tele.Bot) {
	bot.Handle("/start", func(c tele.Context) error {
		return c.Send("Hello!")
	})
}