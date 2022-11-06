package tg

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/taverok/tinyAssistant/internal/app/assistbot/tg/domain"
	"log"
)

type Dispatcher struct {
	handlersMap map[string]func(Request) string
}

func NewDispatcher() Dispatcher {
	return Dispatcher{
		handlersMap: map[string]func(Request) string{},
	}
}

func (it *Dispatcher) RegisterHandler(command string, action func(Request) string) {
	it.handlersMap[command] = action
}

func (it *Dispatcher) Dispatch(m tgbotapi.Message) domain.TgResponse {
	request := NewRequest(m.Chat.ID, m.From.ID, m.Text)

	log.Printf("raw message %+v", m)
	log.Printf("requset %+v", request)

	var handler = it.getHandler(request)
	response := handler(request)

	return domain.TgResponse{
		ChatId: request.TgChatId,
		Text:   response,
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
