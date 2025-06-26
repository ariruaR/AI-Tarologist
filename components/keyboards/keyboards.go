package keyboards

import (
	tele "gopkg.in/telebot.v4"
)

func CreateBuyKeyboard() *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	btnStarCard := menu.Text("Прогноз по звездам")
	btnNotalCard := menu.Text("Нотальная карта")
	btnAnotherBuy := menu.Text("Еще какая то рандомная хрень")
	menu.Reply(
		menu.Row(btnStarCard),
		menu.Row(btnNotalCard),
		menu.Row(btnAnotherBuy),
	)
	return menu
}
