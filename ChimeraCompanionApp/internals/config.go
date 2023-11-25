package internals

import (
	"encoding/json"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Token string `mapstructure:"TOKEN"`

	Admins []int64 `mapstructure:"ADMINS"`

	IAMServiceEndpoint       string `mapstructure:"IAMSERVICE_ENDPOINT"`
	FinancialServiceEndpoint string `mapstructure:"FINANCIALSERVICE_ENDPOINT"`

	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     string `mapstructure:"REDIS_PORT"`
	RedisDB       int    `mapstructure:"REDIS_DB"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
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
		"TOKEN",
		"IAMSERVICE_ENDPOINT",
		"FINANCIALSERVICE_ENDPOINT",
		"REDIS_HOST",
		"REDIS_PORT",
		"REDIS_DB",
		"REDIS_PASSWORD",
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
