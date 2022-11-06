package assistbot

import (
	"github.com/jmoiron/sqlx"
	"github.com/taverok/tinyAssistant/internal/app/assistbot/config"
	"github.com/taverok/tinyAssistant/internal/app/assistbot/service"
	"github.com/taverok/tinyAssistant/internal/app/assistbot/tg"
	"github.com/taverok/tinyAssistant/internal/app/assistbot/tg/handlers"
)

type App struct {
	Config     config.Config
	Dispatcher tg.Dispatcher
	DB         *sqlx.DB
}

func (it *App) Run() {
	it.Config.Init()
	it.registerRoutes()

	tg.Consumer{
		Config:     it.Config,
		Dispatcher: it.Dispatcher,
	}.Listen()
}

func (it *App) registerRoutes() {
	handler := handlers.Handler{
		AlarmService: service.AlarmService{DB: it.DB},
	}

	it.Dispatcher.RegisterHandler("alarm", handler.HandleAlarm)
}
