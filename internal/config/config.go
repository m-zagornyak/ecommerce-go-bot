package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	IsDebug       bool `yaml:"is_debug" env:"ST_BOT_IS_DEBUG" env-default:"false"`
	IsDevelopment bool `yaml:"is_development" env:"ST_BOT_IS_DEVELOPMENT" env-default:"false"`
	Telegram      struct {
		Token string `yaml:"Token" env:"ST_BOT_TELEGRAM_TOKEN" env-default:"5347241093:AAHhBE-pyCpUDQitwmbZG0BNnLsnbJg5u2E"`
	}
	AppConfig AppConfig `yaml:"app"`
}

type AppConfig struct {
	LogLevel string `yaml:"log_level" env:"ST_BOT_LOG_LEVEL" env-default:"error"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}

		if err := cleanenv.ReadEnv(instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Print(help)
			log.Fatal(err)
		}
	})
	return instance
}

/*

func ConfigEnv() {
	err := godotenv.Load("")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	TOKEN := os.Getenv("S3_BUCKET")
}
*/
