package domain

type Alarm struct {
	Cron
	Id       string
	Say      string
	IsActive bool
	UserId   int64
	ChatId   int64
}

func AlarmOf(fields map[string]string, userId int64, chatId int64) (Alarm, error) {
	return Alarm{
		Id:       fields["id"],
		Cron:     Cron{},
		Say:      fields[TOKEN_SAY],
		IsActive: true,
		UserId:   userId,
		ChatId:   chatId,
	}, nil
}
