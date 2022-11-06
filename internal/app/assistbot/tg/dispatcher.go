package tg

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Dispatcher struct {
	Bot
	handlersMap map[string]func(Request) string
}

func NewDispatcher(bot Bot) Dispatcher {
	return Dispatcher{
		Bot:         bot,
		handlersMap: map[string]func(Request) string{},
	}
}

func (it *Dispatcher) ListenAndRoute() {
	bot := it.GetBotApi()
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		request := NewRequest(update.Message.Chat.ID, update.Message.From.ID, update.Message.Text)
		request.ReplyMessageId = update.Message.MessageID
		response := it.Dispatch(request)

		if response.Text == "" {
			continue
		}

		it.Send(response)
	}
}

func (it *Dispatcher) RegisterHandler(command string, action func(Request) string) {
	it.handlersMap[command] = action
}

func (it *Dispatcher) Dispatch(r Request) TgResponse {
	log.Printf("requset %+v", r)

	var handler = it.getHandler(r)
	response := handler(r)

	return TgResponse{
		ChatId:         r.TgChatId,
		Text:           response,
		ReplyMessageId: r.ReplyMessageId,
	}
}

func (it *Dispatcher) getHandler(r Request) func(Request) string {
	handler, ok := it.handlersMap[r.Command]
	if !ok {
		return func(request Request) string {
			return fmt.Sprintf("Error: handler for command %s not found", r.Command)
		}
	}

	return handler
}
