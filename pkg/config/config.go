package config

import "github.com/spf13/viper"

type Config struct {
	TelegramToken string

	Messages Messages
}

type Messages struct {
	Responses
}

type Responses struct {
	Start string
}

func Init() (*Config, error) {
	viper.SetConfigName("configs")
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := parseEnv(&cfg); err != nil {
		return nil, err
	}
}

func parseEnv(cfg *Config) error {
	if err := viper.BindEnv("token"); err != nil {
		return err
	}

	cfg.TelegramToken = viper.GetString("token")
	return nil
}
