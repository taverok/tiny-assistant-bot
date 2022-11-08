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
		"INSERT INTO alarm(id, say, is_active, user_id, chat_id, cron, counter, scheduled_at, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		a.Id, a.Say, a.IsActive, a.UserId, a.ChatId, a.Cron, a.Counter, a.ScheduledAt, a.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (it *AlarmService) UpdateAlarm(a *domain.Alarm) error {
	_, err := it.DB.Exec(
		"UPDATE alarm SET say=?, is_active=?, cron=?, scheduled_at=? WHERE id=?",
		a.Say, a.IsActive, a.Cron, a.ScheduledAt, a.Id)
	if err != nil {
		return err
	}

	log.Printf("updating alarm id: %s", a.Id)
	return nil
}

func (it *AlarmService) RunScheduler() {
	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(1).Minute().Do(func() {
		now := time.Now().UTC()
		log.Printf("alarm scheduler started < %d", now.Unix())
		for _, a := range it.getScheduled(now) {
			it.runAlarm(a, now)
		}
	})

	if err != nil {
		log.Printf("Error scheduling %s", err.Error())
		return
	}

	s.StartAsync()
}

func (it *AlarmService) runAlarm(a domain.Alarm, now time.Time) {
	it.TgBot.Send(tg.TgResponse{
		ChatId: a.ChatId,
		Text:   a.Say,
	})

	a.ScheduleNext(now)
	a.Counter++

	err := it.UpdateAlarm(&a)
	if err != nil {
		log.Println("Error " + err.Error())
	}
}

func (it *AlarmService) getScheduled(now time.Time) []domain.Alarm {
	var aa []domain.Alarm
	err := it.DB.Select(&aa, "SELECT * FROM alarm WHERE scheduled_at <= ? AND is_active=1", now.Unix())
	if err != nil {
		log.Println(err)
		return nil
	}

	return aa
}
