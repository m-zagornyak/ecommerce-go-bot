package main

import (
	"flag"
	"github.com/m-zagornyak/ecommerce-go-bot/internal"
	"github.com/m-zagornyak/ecommerce-go-bot/internal/config"
	"github.com/m-zagornyak/ecommerce-go-bot/pkg/logging"
	"log"
)

var cfgPath string

func init() {
	flag.StringVar(&cfgPath, "config", "configs/prod.yml", "using file path")
}

func main() {
	log.Print("config initializing")
	cfg := config.GetConfig(cfgPath)

	log.Print("logger initializing")
	logging.Init(cfg.AppConfig.LogLevel)
	logger := logging.GetLogger()

	logger.Println("Creating Application")
	app, err := internal.NewApp(logger, cfg)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("Running Application")
	app.Run()
}
