package config

import "os"

type Config struct {
	TgKey  string
	DbPath string
}

func NewConfig() Config {
	config := Config{}
	config.TgKey = os.Getenv("TINY_ASSISTANT_TG_KEY")

	dbPath, ok := os.LookupEnv("TINY_ASSISTANT_DB_PATH")
	if !ok {
		dbPath = "db.sql"
	}
	config.DbPath = dbPath

	return config
}
