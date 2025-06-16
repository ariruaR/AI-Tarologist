package keyboards

import (
	"github.com/mymmrac/telego"
	ti "github.com/mymmrac/telego/telegoutil"
)

func Createkeyboard() telego.ReplyKeyboardMarkup {
	keyboard := ti.Keyboard(
		ti.KeyboardRow(
			ti.KeyboardButton("Прогноз по звездам"),
			ti.KeyboardButton("Нотальная карта"),
			ti.KeyboardButton("еще какая то очень дорогая хрень"),
		),
		ti.KeyboardRow(
			ti.KeyboardButton("оплатить все это"),
		),
	).WithResizeKeyboard().WithInputFieldPlaceholder("Выбрать что-то одно")
	return *keyboard
}
