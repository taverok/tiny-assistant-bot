package domain

import (
	"github.com/google/uuid"
)

type Alarm struct {
	Cron
	Id       string
	Message  string
	IsActive bool
	UserId   int64
	ChatId   int64
}

func AlarmOf(fields map[string]string, userId int64, chatId int64) Alarm {
	return Alarm{
		Id:       uuid.New().String(),
		Cron:     Cron{},
		Message:  fields[TOKEN_SAY],
		IsActive: true,
		UserId:   userId,
		ChatId:   chatId,
	}
}
