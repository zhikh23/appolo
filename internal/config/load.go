package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

func MustLoad() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("CONFIG_PATH is empty")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("CONFIG_PATH is empty: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
