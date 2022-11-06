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
	TgBot      tg.Bot
}

func NewApp() App {
	conf := config.NewConfig()
	bot := tg.Bot{
		Config: conf,
	}

	app := App{
		Config:     conf,
		Dispatcher: tg.NewDispatcher(bot),
		DB:         sqlx.MustConnect("sqlite3", conf.DbPath),
		TgBot:      bot,
	}

	app.registerRoutes()

	return app
}

func (it *App) Run() {

	it.Dispatcher.ListenAndRoute()
}

func (it *App) registerRoutes() {
	handler := handlers.Handler{
		AlarmService: service.AlarmService{DB: it.DB, TgBot: it.TgBot},
	}

	it.Dispatcher.RegisterHandler("alarm", handler.HandleAlarm)

	handler.AlarmService.RunScheduler()
}
