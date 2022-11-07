package config

import "os"

type Config struct {
	TgKey  string
	DbPath string
}

func NewConfig() Config {
	config := Config{}
	config.TgKey = getEnv("TINY_ASSISTANT_TG_KEY", "")
	config.DbPath = getEnv("TINY_ASSISTANT_DB_PATH", "db.sqlite")

	return config
}

func getEnv(key string, def string) string {
	v, ok := os.LookupEnv(key)
	if ok {
		return v
	}
	return def
}
