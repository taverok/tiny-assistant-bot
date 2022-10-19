package domain

type Request struct {
	Path     string
	Fields   map[string]string
	TgChatId int64
	TgUserId int64
}
