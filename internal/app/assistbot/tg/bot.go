package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/taverok/tinyAssistant/internal/app/assistbot/config"
	"log"
)

type Bot struct {
	Config config.Config
	botApi *tgbotapi.BotAPI
}

func (it *Bot) Send(r TgResponse) {
	bot := it.GetBotApi()
	msg := tgbotapi.NewMessage(r.ChatId, r.Text)
	msg.DisableWebPagePreview = true

	if r.ReplyMessageId != 0 {
		msg.ReplyToMessageID = r.ReplyMessageId
	}

	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func (it *Bot) GetBotApi() *tgbotapi.BotAPI {
	if it.botApi != nil {
		return it.botApi
	}

	bot, err := tgbotapi.NewBotAPI(it.Config.TgKey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	return bot
}
