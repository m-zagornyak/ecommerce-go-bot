package internal

import (
	"github.com/m-zagornyak/ecommerce-go-bot/internal/config"
	"github.com/m-zagornyak/ecommerce-go-bot/pkg/logging"
	tele "gopkg.in/telebot.v3"
	"log"
	"net/http"
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
	b, err := tele.NewBot(tele.Settings{
		Token:  a.cfg.Telegram.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	var (
		menu     = &tele.ReplyMarkup{ResizeKeyboard: true}
		selector = &tele.ReplyMarkup{}

		btnHelp         = menu.Text("ℹ Поддержка")
		btnCategory     = menu.Text("Все категории")
		btnRecyclingBin = menu.Text("🗑 Корзина")

		//	btnPrev = selector.Data("⬅", "prev", "help")
		//	btnNext = selector.Data("➡", "next", "me")
	)

	menu.Reply(
		menu.Row(btnHelp, btnRecyclingBin),
		//menu.Row(btnRecyclingBin),
		menu.Row(btnCategory),
	)
	selector.Inline(
		selector.Row(btnHelp, btnRecyclingBin),
		selector.Row(btnCategory),
	)

	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Hello!", menu)
	})

	b.Handle(&btnHelp, func(c tele.Context) error {
		return c.Send("Here is some help: ...")
	})

	b.Handle(&btnRecyclingBin, func(c tele.Context) error {
		return c.Respond()
	})

	b.Start()
}
