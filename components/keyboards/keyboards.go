package keyboards

import (
	tele "gopkg.in/telebot.v4"
)

var menu = &tele.ReplyMarkup{ResizeKeyboard: true}
var BtnStarCard = menu.Text("Прогноз по звездам")
var BtnNotalCard = menu.Text("Нотальная карта")
var BtnAnotherBuy = menu.Text("Еще какая то рандомная хрень")

func CreateBuyKeyboard() *tele.ReplyMarkup {
	menu.Reply(
		menu.Row(BtnStarCard),
		menu.Row(BtnNotalCard),
		menu.Row(BtnAnotherBuy),
	)
	return menu
}
