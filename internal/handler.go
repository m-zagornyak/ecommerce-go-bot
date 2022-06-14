package internal

import tele "gopkg.in/telebot.v3"


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