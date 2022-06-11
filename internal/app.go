package internal

import (
	"github.com/m-zagornyak/ecommerce-go-bot/internal/config"
	"github.com/m-zagornyak/ecommerce-go-bot/pkg/logging"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"net/http"
	"os"
	"time"
)

type app struct {
	cfg        *config.Config
	logger     *logging.Logger
	httpServer *http.Server
}

type App interface {
	Run()
}

func NewApp(logger *logging.Logger, cfg *config.Config) (App, error) {
	logger.Println("")
	return &app{
		cfg:    cfg,
		logger: logger,
	}, nil
}

func (a *app) Run() {
	a.startBot()
}

func (a *app) startBot() {
	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("5347241093:AAHhBE-pyCpUDQitwmbZG0BNnLsnbJg5u2E"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "hello world")
	})

	b.Start()
}
