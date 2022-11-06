package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/taverok/tinyAssistant/internal/app/assistbot"
)

func main() {
	app := assistbot.NewApp()
	app.Run()
}
