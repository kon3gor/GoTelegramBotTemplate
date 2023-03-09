package main

import (
	"com/example/gobot/internal/config"
	"com/example/gobot/internal/telegram"
)

func main() {
	if err := config.ReadConfig(); err != nil {
		panic(err)
	}

	if err := telegram.InitTelegramBot(); err != nil {
		panic(err)
	}
	telegram.StartPolling(true)
}
