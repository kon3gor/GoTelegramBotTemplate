package main

import (
	"com/example/gobot/internal/config"
	"com/example/gobot/internal/logs"
	"com/example/gobot/internal/telegram"
	"log"
)

func main() {
	f, err := logs.Init()
	if err != nil {
		panic(err)
	}
	defer f()
	if err := config.ReadConfig(); err != nil {
		log.Fatalln(err)
	}

	if err := telegram.InitTelegramBot(); err != nil {
		log.Fatalln(err)
	}
	telegram.StartPolling(true)
}
