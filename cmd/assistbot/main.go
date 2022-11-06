package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/taverok/tinyAssistant/internal/app/assistbot"
	"github.com/taverok/tinyAssistant/internal/app/assistbot/tg"
	"os"
)

func main() {
	dbPath, ok := os.LookupEnv("TINY_ASSISTANT_DB_PATH")
	if !ok {
		dbPath = "db.sql"
	}

	app := assistbot.App{
		Dispatcher: tg.NewDispatcher(),
		DB:         sqlx.MustConnect("sqlite3", dbPath),
	}
	app.Run()
}
