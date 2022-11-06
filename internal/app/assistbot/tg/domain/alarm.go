package domain

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"github.com/taverok/tinyAssistant/internal/app/assistbot/utils"
	"time"
)

type Alarm struct {
	Id          string `db:"id"`
	Say         string `db:"say"`
	IsActive    bool   `db:"is_active"`
	UserId      int64  `db:"user_id"`
	ChatId      int64  `db:"chat_id"`
	Cron        string `db:"cron"`
	ScheduledAt int64  `db:"scheduled_at"`
	CreatedAt   int64  `db:"created_at"`
	Counter     int    `db:"counter"`
}

func (it *Alarm) ScheduleNext(now time.Time) {
	scheduleUTC := cronexpr.MustParse(it.Cron).Next(now)
	it.ScheduledAt = scheduleUTC.Unix()
}

func (it *Alarm) String() string {
	return fmt.Sprintf(`alarm
id %s
on %s
say %s`, it.Id, it.Cron, it.Say)
}

func AlarmOf(fields map[string]string, userId int64, chatId int64) (Alarm, error) {
	cron := fields["on"]
	if len(cron) == 0 {
		cron = utils.DEFAULT_CRON
	}
	ok := utils.CronRegex.MatchString(cron)
	if !ok {
		return Alarm{}, utils.ErrorWrongCronRegex
	}

	now := time.Now().UTC()

	alarm := Alarm{
		Id:        fields["id"],
		Cron:      cron,
		Say:       fields[TOKEN_SAY],
		IsActive:  true,
		UserId:    userId,
		ChatId:    chatId,
		CreatedAt: now.Unix(),
		Counter:   0,
	}

	alarm.ScheduleNext(now)

	return alarm, nil
}
