package Telegram

import (
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func init () {

}

func BotInit(botToken string) {
	//bot, err := tgbotapi.NewBotAPI(botToken)
	//781061948:AAHhMNvHv2RN0uIarC8DIQP9F6ShYhP0CY8
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}