package config

import (
	"fmt"
	"io/ioutil"
	"os"

	toml "github.com/pelletier/go-toml/v2"
)

type logs struct {
	write_logs bool
	logs_path  string
}

type telegram struct {
	token string
}

type bot struct {
	log_unknown bool
}

type Config struct {
	logs     logs
	telegram telegram
	bot      bot
}

var config Config

func ReadConfig() error {
	f, err := os.Open("config.toml")
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	toml.Unmarshal(b, &config)
	fmt.Println(config.logs.write_logs)

	return nil
}

