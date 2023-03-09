package process

import (
	"com/example/gobot/internal/appcontext"
	"log"
)

func ProcessCommand(context *appcontext.Context) {
	command := context.RawUpdate.Message.Command()

	switch command {
	default:
		unknownCommand(context, command)
	}
}

func unknownCommand(context *appcontext.Context, command string) {
	log.Println("Recieved unknown  command: ", command)
	context.TextAnswer("Unknown command")
}
