package domain

type Alarm struct {
	Cron
	Id       string `db:"id"`
	Say      string `db:"say"`
	IsActive bool   `db:"is_active"`
	UserId   int64  `db:"user_id"`
	ChatId   int64  `db:"chat_id"`
}

func AlarmOf(fields map[string]string, userId int64, chatId int64) (Alarm, error) {
	cron, err := CronFromString(fields["on"])
	if err != nil {
		return Alarm{}, err
	}

	return Alarm{
		Id:       fields["id"],
		Cron:     cron,
		Say:      fields[TOKEN_SAY],
		IsActive: true,
		UserId:   userId,
		ChatId:   chatId,
	}, nil
}
