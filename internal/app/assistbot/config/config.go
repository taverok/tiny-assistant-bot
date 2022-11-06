package config

import "os"

type Config struct {
	TgKey string
}

func (it *Config) Init() {
	it.TgKey = os.Getenv("TINY_ASSISTANT_TG_KEY")
}
