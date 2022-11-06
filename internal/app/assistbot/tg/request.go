package tg

import (
	"strings"
)

type Request struct {
	Command        string
	Fields         map[string]string
	TgChatId       int64
	TgUserId       int64
	ReplyMessageId int
}

func NewRequest(chatId int64, userId int64, message string) Request {
	command, fields := parseMessage(message)

	return Request{
		Command:  command,
		Fields:   fields,
		TgChatId: chatId,
		TgUserId: userId,
	}
}

// parseMessage return command and map of fields
func parseMessage(m string) (string, map[string]string) {
	lines := strings.Split(m, "\n")
	if len(lines) == 0 {
		return "", map[string]string{}
	}

	firstLine := strings.TrimSpace(lines[0])
	bodyLines := lines[1:]
	path := parseCommand(strings.Split(firstLine, " "))
	fields := parseFields(bodyLines)

	return path, fields
}

func parseFields(lines []string) map[string]string {
	result := map[string]string{}
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		if len(tokens) == 0 {
			continue
		}

		for i, token := range tokens {
			tokens[i] = normalizeToken(token)
		}

		result[tokens[0]] = strings.Join(tokens[1:], " ")
	}

	return result
}

func parseCommand(tokens []string) string {
	return strings.ToLower(tokens[0])
}

func normalizeToken(token string) string {
	t := strings.TrimSpace(token)
	t = strings.ToLower(t)
	return t
}
