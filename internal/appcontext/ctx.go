package appcontext

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Context struct {
	UserName  string
	ChatID    int64
	RawUpdate tgbotapi.Update
	Args      string

	bot *tgbotapi.BotAPI
}

func FromCommand(update tgbotapi.Update, bot *tgbotapi.BotAPI) *Context {
	chatId := update.Message.Chat.ID
	username := update.Message.From.UserName
	args := update.Message.CommandArguments()
	return &Context{username, chatId, update, args, bot}
}

func FromCallback(update tgbotapi.Update, bot *tgbotapi.BotAPI) *Context {
	chatId := update.CallbackQuery.Message.Chat.ID
	username := update.CallbackQuery.Message.From.UserName
	return &Context{username, chatId, update, "", bot}
}
