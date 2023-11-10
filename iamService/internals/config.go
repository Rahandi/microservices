package internals

import (
	"encoding/json"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"PORT"`

	DatabaseName     string `mapstructure:"DB_NAME"`
	DatabaseHost     string `mapstructure:"DB_HOST"`
	DatabasePort     string `mapstructure:"DB_PORT"`
	DatabaseUsername string `mapstructure:"DB_USERNAME"`
	DatabasePassword string `mapstructure:"DB_PASSWORD"`

	JwtSecret         string `mapstructure:"JWT_SECRET"`
	JwtRefreshSecret  string `mapstructure:"JWT_REFRESH_SECRET"`
	JwtExpires        string `mapstructure:"JWT_EXPIRES"`         // Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	JwtRefreshExpires string `mapstructure:"JWT_REFRESH_EXPIRES"` // Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
}

func NewConfig() *Config {
	config := Config{}

	file, err := os.Open("config.json")
	if err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
	} else {
		defer file.Close()
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&config)
		if err != nil {
			panic(err)
		}
	}

	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	// Bind each key in the Config struct to a corresponding environment variable
	keys := []string{
		"PORT",
		"DB_NAME",
		"DB_HOST",
		"DB_PORT",
		"DB_USERNAME",
		"DB_PASSWORD",
		"JWT_SECRET",
		"JWT_REFRESH_SECRET",
		"JWT_EXPIRES",
		"JWT_REFRESH_EXPIRES",
	}
	for _, key := range keys {
		viper.BindEnv(key)
	}

	_ = viper.ReadInConfig() // Ignore error if .env file does not exist
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	return &config
}
