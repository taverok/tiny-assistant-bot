package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/taverok/tinyAssistant/app/assistant/tg/domain"
	"log"
	"strings"
)

type Dispatcher struct {
	Handler Handler
}

func (it *Dispatcher) Dispatch(m tgbotapi.Message) domain.TgResponse {
	chatId := m.Chat.ID
	userId := m.From.ID
	path, fields := it.parseMessage(m)

	log.Printf("raw message %+v", m)
	log.Printf("path %s body %+v", path, fields)

	response := it.Handler.Handle(domain.Request{
		Path:     path,
		Fields:   fields,
		TgChatId: chatId,
		TgUserId: userId,
	})

	return domain.TgResponse{
		ChatId: chatId,
		Text:   response,
	}
}

// parseMessage return path and map of fields
func (it *Dispatcher) parseMessage(m tgbotapi.Message) (string, map[string]string) {
	lines := strings.Split(m.Text, "\n")
	if len(lines) == 0 {
		return "", map[string]string{}
	}

	firstLine := strings.TrimSpace(lines[0])
	bodyLines := lines[1:]
	path := tokensToPath(strings.Split(firstLine, " "))
	fields := linesToFieldsMap(bodyLines)

	return path, fields
}

func linesToFieldsMap(lines []string) map[string]string {
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

func tokensToPath(tokens []string) string {
	var method string
	var pathTokens []string
	for i, token := range tokens {
		if i == 0 {
			method = strings.ToUpper(token)
			continue
		}
		t := normalizeToken(token)
		pathTokens = append(pathTokens, t)
	}

	return method + ":" + strings.Join(pathTokens, "-")
}

func normalizeToken(token string) string {
	t := strings.TrimSpace(token)
	t = strings.ToLower(t)
	return t
}
