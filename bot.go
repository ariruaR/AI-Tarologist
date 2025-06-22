package main

import (
	"context"
	"fmt"
	"os"

	"bot/components/chatgpt"
	"bot/components/keyboards"
	"bot/components/payment"
	configReader "bot/config"

	"github.com/mymmrac/telego"
	ti "github.com/mymmrac/telego/telegoutil"
)

// Local Imports

const StartText string = `Привет, я ИИ-Таролог, составлю натальную карту и погадаю на таро, а так же расскажу Вам Вашу судьбу по знаку зодиака\n
Всего за 70$ ты узнаешь будущее себя и своей семьи
Для приобретения услуги - пиши /buy
`

const BuyText string = `
Правильный выбор! 
Вот тебе стоимости: 
Прогноз по звездам: 70.0 $ или 7000 ⭐
Нотальная карта: 250.0 $ или 25000 ⭐
Еще какая то очень дорогая хрень: 2500.0$ или 250000 ⭐
`

func main() {
	bot_token := configReader.Readconfig().BOTTOKEN

	ctx := context.Background()

	bot, err := telego.NewBot(bot_token, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	updates, _ := bot.UpdatesViaLongPolling(context.Background(), nil)
	for update := range updates {
		chatId := update.Message.Chat.ID
		if update.Message.Text == "/start" {
			bot.SendMessage(ctx,
				ti.Message(
					ti.ID(chatId),
					StartText,
				),
			)
		}
		if update.Message.Text == "/buy" {
			buykeyboard := keyboards.Createkeyboard()
			msg := ti.Message(
				ti.ID(chatId), BuyText,
			).WithReplyMarkup(&buykeyboard).WithProtectContent()
			bot.SendMessage(ctx, msg)
		}
		if update.Message.Text == "Прогноз по звездам" {
			payment.SendInvoice(
				bot,
				ctx,
				chatId,
				"Прогноз по звездам",
				"Оплата стоимости прогноза по звездам",
				7000,
			)
			if update.Message.SuccessfulPayment != nil {
				responseMsg := fmt.Sprintf("Составь прогноз на неделю по звездам для %s, учитывая все особенности тарологических прогнозов", update.Message.From.FirstName)
				response := chatgpt.RequestOpenAi(responseMsg)
				bot.SendMessage(ctx,
					ti.Message(
						ti.ID(chatId),
						response,
					),
				)
			}

		}
		if update.Message.Text == "Нотальная карта" {
			payment.SendInvoice(
				bot,
				ctx,
				chatId,
				"Нотальная карта",
				"Оплата стоимости Нотальная карта",
				25000,
			)
			if update.Message.SuccessfulPayment != nil {
				responseMsg := fmt.Sprintf("Составь нотальная карта для %s, учитывая все особенности тарологических прогнозов", update.Message.From.FirstName)
				response := chatgpt.RequestOpenAi(responseMsg)
				bot.SendMessage(ctx,
					ti.Message(
						ti.ID(chatId),
						response,
					),
				)
			}

		}
		if update.Message.Text == "еще какая то очень дорогая хрень" {
			payment.SendInvoice(
				bot,
				ctx,
				chatId,
				"еще какая то очень дорогая хрень",
				"Оплата стоимости еще какая то очень дорогая хрень",
				250000,
			)
			if update.Message.SuccessfulPayment != nil {
				responseMsg := fmt.Sprintf("Составь прогноз на неделю по звездам для %s, учитывая все особенности тарологических прогнозов", update.Message.From.FirstName)
				response := chatgpt.RequestOpenAi(responseMsg)
				bot.SendMessage(ctx,
					ti.Message(
						ti.ID(chatId),
						response,
					),
				)
			}

		}
		if update.Message.Text == "оплатить все это" {
			payment.SendInvoice(
				bot,
				ctx,
				chatId,
				"еще какая то очень дорогая хрень",
				"Оплата стоимости еще какая то очень дорогая хрень",
				270000,
			)
			if update.Message.SuccessfulPayment != nil {
				// responseMsg := fmt.Sprintf("Составь прогноз на неделю по звездам для %s, учитывая все особенности тарологических прогнозов", update.Message.From.FirstName)
				// response := chatgpt.RequestOpenAi(responseMsg)
				// bot.SendMessage(ctx,
				// 	ti.Message(
				// 		ti.ID(chatId),
				// 		response,
				// 	),
				// )
				bot.SendMessage(ctx, ti.Message(ti.ID(chatId), "Запрос создан, ожидайте ответа от менеджера"))
			}
		}
	}
}
