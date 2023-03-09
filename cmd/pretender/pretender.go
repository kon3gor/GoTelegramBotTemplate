package main

import "com/example/gobot/internal/telegram"

func main() {
	if err := telegram.InitTelegramBot(); err != nil {
		panic(err)
	}
	telegram.StartPolling(true)
}
