package main

import (
	"github.com/m-zagornyak/ecommerce-go-bot/internal/config"
	"github.com/m-zagornyak/ecommerce-go-bot/internal/telegram"
	"github.com/m-zagornyak/ecommerce-go-bot/pkg/logging"
	"log"
)

/*
var cfgPath string

func init() {
	flag.StringVar(&cfgPath, "config", "configs/prod.yml", "using file path")
}
*/
func main() {
	log.Print("config initializing")
	cfg := config.GetConfig()

	log.Print("logger initializing")
	logging.Init(cfg.AppConfig.LogLevel)
	logger := logging.GetLogger()

	logger.Println("Creating Application")
	app, err := telegram.NewApp(logger, cfg)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("Running Application")
	app.Run()
}
