package payment

import (
	tele "gopkg.in/telebot.v4"
)

func SendInvoice(
	ctx tele.Context,
	title string,
	description string,
	amount uint64,
) error {
	invoice := tele.Invoice{
		Title:       title,
		Description: description,
		Payload:     "unique_payment",
		Token:       "",
		Currency:    "XTR",
		Prices: []tele.Price{
			{
				Label:  "Оплата запроса",
				Amount: int(amount),
			},
		},
	}
	return ctx.Send(invoice)
}
