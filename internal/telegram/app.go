package telegram

import (
	"github.com/m-zagornyak/ecommerce-go-bot/internal/config"
	"github.com/m-zagornyak/ecommerce-go-bot/pkg/logging"
	tele "gopkg.in/telebot.v3"
	"log"
	"time"

	"net/http"
	//"time"
)

type App struct {
	Cfg        *config.Config
	Logger     *logging.Logger
	HttpServer *http.Server
	Bot 	     *tele.Bot
}

type BaseApp interface {
	Run()
}

func NewApp(logger *logging.Logger, cfg *config.Config) (BaseApp, error) {
	logger.Println("")
	return &App{
		Cfg:    cfg,
		Logger: logger,

	}, nil
}

func (a *App) Run() {
	a.startBot()
}

func (a *App) NewBot() *tele.Bot{

	b, err := tele.NewBot(tele.Settings{
		Token:  a.cfg.Telegram.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return b
	}

/*
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
*/