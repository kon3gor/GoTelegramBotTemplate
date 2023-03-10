package logs

import (
	"log"
	"os"
)

type deferable func() error

func Init() (deferable, error) {
	f, err := os.Create("./log.log")
	if err != nil {
		return nil, err
	}

	log.SetOutput(f)
	log.SetFlags(log.Lshortfile)

	return f.Close, nil
}
