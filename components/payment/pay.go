package payment

import (
	"context"
	"log"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func SendInvoice(
	bot *telego.Bot,
	ctx context.Context,
	chatID int64,
	title string,
	description string,
	amount uint64,
) {
	invoice := telego.SendInvoiceParams{
		ChatID:        tu.ID(chatID),
		Title:         title,
		Description:   description,
		Payload:       "unique_payment",
		ProviderToken: "",
		Currency:      "XTR",
		Prices: []telego.LabeledPrice{
			{
				Label:  "Оплата запроса",
				Amount: int(amount),
			},
		},
	}
	_, err := bot.SendInvoice(ctx, &invoice)
	if err != nil {
		log.Printf("Ошибка отправки чека оплаты, err: %s", err)
	}
}
