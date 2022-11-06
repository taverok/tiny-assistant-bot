package service

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/taverok/tinyAssistant/internal/app/assistbot/tg/domain"
)

type AlarmService struct {
	DB *sqlx.DB
}

func (it *AlarmService) CreateAlarm(alarm *domain.Alarm) error {
	alarm.Id = uuid.New().String()

	_, err := it.DB.Exec("INSERT INTO alarm(id, say, is_active, user_id, chat_id) VALUES (?, ?, ?, ?, ?)",
		alarm.Id, alarm.Say, alarm.IsActive, alarm.UserId, alarm.ChatId,
	)
	if err != nil {
		return err
	}

	return nil
}
