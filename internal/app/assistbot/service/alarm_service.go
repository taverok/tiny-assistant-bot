package service

import (
	"github.com/go-co-op/gocron"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/taverok/tinyAssistant/internal/app/assistbot/tg"
	"github.com/taverok/tinyAssistant/internal/app/assistbot/tg/domain"
	"log"
	"time"
)

type AlarmService struct {
	DB    *sqlx.DB
	TgBot tg.Bot
}

func (it *AlarmService) CreateAlarm(a *domain.Alarm) error {
	a.Id = uuid.New().String()

	_, err := it.DB.Exec(
		"INSERT INTO alarm(id, say, is_active, user_id, chat_id, `minute`, `hour`, day_month, `month`, day_week) "+
			"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		a.Id, a.Say, a.IsActive, a.UserId, a.ChatId, a.Minute, a.Hour, a.DayMonth, a.Month, a.DayWeek,
	)
	if err != nil {
		return err
	}

	return nil
}

func (it *AlarmService) RunScheduler() {
	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(1).Minute().Do(func() {
		log.Printf("alarm scheduler started")
		for _, a := range it.getScheduled() {
			it.runAlarm(a)
		}
	})

	if err != nil {
		log.Printf("Error scheduling %s", err.Error())
		return
	}

	s.StartAsync()
}

func (it *AlarmService) runAlarm(a domain.Alarm) {
	it.TgBot.Send(tg.TgResponse{
		ChatId: a.ChatId,
		Text:   a.Say,
	})
}

func (it *AlarmService) getScheduled() []domain.Alarm {
	now := time.Now()

	var aa []domain.Alarm
	err := it.DB.Select(&aa, "SELECT * FROM alarm WHERE minute in ('*', ?)", now.Minute())
	if err != nil {
		log.Println(err)
		return nil
	}

	return aa
}
