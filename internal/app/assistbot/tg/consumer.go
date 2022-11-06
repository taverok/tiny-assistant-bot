package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/taverok/tinyAssistant/internal/app/assistbot/config"
	"log"
)

type Consumer struct {
	Config     config.Config
	Dispatcher Dispatcher
}

func (it Consumer) Listen() {
	bot, err := tgbotapi.NewBotAPI(it.Config.TgKey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {

			response := it.Dispatcher.Dispatch(*update.Message)
			if response.Text == "" {
				continue
			}

			msg := tgbotapi.NewMessage(response.ChatId, response.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			_, err := bot.Send(msg)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
