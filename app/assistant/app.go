package assistant

import (
	config2 "github.com/taverok/tinyAssistant/app/assistant/config"
	"github.com/taverok/tinyAssistant/app/assistant/tg"
)

func Run() {
	var config config2.Config
	config.Init()

	tg.Consumer{
		Config: config,
		Dispatcher: tg.Dispatcher{
			Handler: tg.Handler{},
		},
	}.Listen()
}
