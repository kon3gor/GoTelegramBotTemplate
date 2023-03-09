package process

import (
	"com/example/gobot/internal/appcontext"
	"log"
	"strings"
)

func ProcessCallback(context *appcontext.Context) {
	data := context.RawUpdate.CallbackQuery.Data
	callback, _, _ := strings.Cut(data, ":")
	switch callback {
	default:
		unknownCallback(context, callback)
	}
}

func unknownCallback(context *appcontext.Context, callback string) {
	log.Println("Recieved unknown callback: ", callback)
	context.TextAnswer("Unknown callback")
}
