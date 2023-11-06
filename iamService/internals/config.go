package internals

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port string

	Database struct {
		Host     string
		Port     string
		Username string
		Password string
		Database string
	}

	Jwt struct {
		Secret string
	}
}

func NewConfig() *Config {
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}

	return &config
}
