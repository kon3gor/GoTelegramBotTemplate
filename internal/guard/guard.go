package guard

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type GuardErr struct {
	msg string
}

func NewGuardErr(msg string) *GuardErr {
	return &GuardErr{msg}
}

func (self *GuardErr) Error() string {
	return self.msg
}

type GuardFunc func(tgbotapi.Update) *GuardErr
