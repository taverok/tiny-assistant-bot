package tg

import (
	"fmt"
	"github.com/taverok/tinyAssistant/app/assistant/tg/domain"
)

type Handler struct {
}

func (it *Handler) Handle(r domain.Request) string {
	switch r.Path {
	case "NEW:alarm":
		request := domain.AlarmOf(r.Fields, r.TgUserId, r.TgChatId)
		return it.HandleCreateAlarm(&request)
	default:
		return fmt.Sprintf("handler for %s not found", r.Path)
	}
}

func (it *Handler) HandleCreateAlarm(alarm *domain.Alarm) string {
	return fmt.Sprintf("new alarm created: %+v", alarm)
}
