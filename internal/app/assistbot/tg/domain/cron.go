package domain

import (
	"github.com/taverok/tinyAssistant/internal/app/assistbot/utils"
	"strings"
)

type Cron struct {
	Minute   string `db:"minute"`
	Hour     string `db:"hour"`
	DayMonth string `db:"day_month"`
	Month    string `db:"month"`
	DayWeek  string `db:"day_week"`
}

func CronFromString(s string) (Cron, error) {
	s = utils.MultiSpaceRegex.ReplaceAllString(s, " ")

	if !utils.CronRegex.MatchString(s) {
		return Cron{}, utils.ErrorWrongCronRegex
	}

	ss := strings.Split(s, " ")

	return Cron{
		Minute:   ss[0],
		Hour:     ss[1],
		DayMonth: ss[2],
		Month:    ss[3],
		DayWeek:  ss[4],
	}, nil
}
