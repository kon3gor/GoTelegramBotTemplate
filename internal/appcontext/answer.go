package appcontext

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (self *Context) CustomAnswer(msg tgbotapi.Chattable) {
	self.sendChattable(msg)
}

func (self *Context) TextAnswer(text string) {
	chatId := self.ChatID
	msg := tgbotapi.NewMessage(chatId, text)
	self.sendChattable(msg)
}

func (self *Context) MardownAnswer(text string) {
	chatId := self.ChatID
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ParseMode = "MarkdownV2"
	self.sendChattable(msg)
}

func (self *Context) StickerAnswer(sticker string) {
	chatId := self.ChatID
	msg := tgbotapi.NewSticker(chatId, tgbotapi.FileID(sticker))
	self.sendChattable(msg)
}

func (self *Context) sendChattable(msg tgbotapi.Chattable) {
	_, err := self.bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
