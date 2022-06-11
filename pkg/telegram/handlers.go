package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const commandStart = "start"

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Команда не найдена ")

	switch message.Command() {
	case commandStart:
		msg.Text = "Привет"

		_, err := b.bot.Send(msg)
		return err
	default:
		_, err := b.bot.Send(msg)
		return err
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	// msg.ReplyToMessageID = update.Message.MessageID

	b.bot.Send(msg)
}

func (b *Bot) StartKeyboard(message *tgbotapi.Message) {

	msg := tgbotapi.NewMessage(message.Chat.ID, "Неверно")

	switch message.Text {
	case "open":
		msg.ReplyMarkup = numericKeyboard

	case "close":
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	}

	if _, err := b.bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
