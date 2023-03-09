package telegram

import (
	"com/example/gobot/internal/appcontext"
	"com/example/gobot/internal/process"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI

func InitTelegramBot() error {
	token := os.Getenv("TELEGRAM_TOKEN")
	tgbot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	bot = tgbot

	return nil
}

func StartPolling(debug bool) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	bot.Debug = debug

	for update := range updates {
		ProcessUpdate(update)
	}
}

func ProcessUpdate(update tgbotapi.Update) {
	if update.Message != nil && update.Message.IsCommand() {
		context := appcontext.FromCommand(update, bot)
		process.ProcessCommand(context)
	} else if update.CallbackQuery != nil {
		context := appcontext.FromCallback(update, bot)
		processCallback(context)
	}
}

func processCallback(context *appcontext.Context) {
	callbackQueryId := context.RawUpdate.CallbackQuery.ID
	callbackQueryData := context.RawUpdate.CallbackQuery.Data
	callback := tgbotapi.NewCallback(callbackQueryId, callbackQueryData)
	if _, err := bot.Request(callback); err != nil {
		panic(err)
	}
	process.ProcessCallback(context)
}
