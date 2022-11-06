package handlers

import (
	"fmt"
	"github.com/taverok/tinyAssistant/internal/app/assistbot/service"
	"github.com/taverok/tinyAssistant/internal/app/assistbot/tg"
	"github.com/taverok/tinyAssistant/internal/app/assistbot/tg/domain"
	"log"
)

type Handler struct {
	service.AlarmService
}

func (it *Handler) HandleAlarm(r tg.Request) string {
	alarm, err := domain.AlarmOf(r.Fields, r.TgUserId, r.TgChatId)
	if err != nil {
		return err.Error()
	}

	switch {
	case alarm.Id == "":
		return it.createAlarm(&alarm)
	default:
		return it.updateAlarm(&alarm)
	}
}

func (it *Handler) createAlarm(alarm *domain.Alarm) string {
	err := it.CreateAlarm(alarm)
	if err != nil {
		log.Printf(err.Error())
		return ""
	}

	return fmt.Sprintf("new alarm created: %+v", alarm)
}

func (it *Handler) updateAlarm(alarm *domain.Alarm) string {
	return fmt.Sprintf("new alarm updated: %+v", alarm)
}
