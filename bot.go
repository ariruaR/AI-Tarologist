package main

import (
	"fmt"
	"os"
	"time"

	"bot/components/keyboards"
	"bot/components/payment"
	configReader "bot/config"

	tele "gopkg.in/telebot.v4"
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
Нотальная карта: 85.0 $ или 8500 ⭐
Еще какая то очень дорогая хрень: 100.0$ или 10000 ⭐
`

func main() {
	bot_token := configReader.Readconfig().BOTTOKEN
	pref := tele.Settings{
		Token:  bot_token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	menu := keyboards.CreateBuyKeyboard()
	bot.Handle(tele.OnCheckout, func(ctx tele.Context) error {
		return ctx.Accept()
	})
	bot.Handle("/refund", func(ctx tele.Context) error {
		return ctx.Send("Эта Функция находится в разработке")
	})
	bot.Handle(tele.OnPayment, func(ctx tele.Context) error {
		user := ctx.Sender()
		text := fmt.Sprintf("%s, Оплата прошла успешно", user.Username)
		return ctx.Send(text)
	})
	bot.Handle("/start", func(ctx tele.Context) error {
		return ctx.Send(StartText)
	})
	bot.Handle("/buy", func(ctx tele.Context) error {
		return ctx.Send(BuyText, menu)
	})

	bot.Handle(&keyboards.BtnStarCard, func(ctx tele.Context) error {
		invoice := payment.CreatePayInvoice(ctx, "Прогноз по звездам", "Оплата услуги", 7000)
		return ctx.Send(invoice)
	})
	bot.Handle(&keyboards.BtnAnotherBuy, func(ctx tele.Context) error {
		invoice := payment.CreatePayInvoice(ctx, "Еще какая то штука", "оплата услуги", 10000)
		return ctx.Send(invoice)
	})
	bot.Handle(&keyboards.BtnNotalCard, func(ctx tele.Context) error {
		invoice := payment.CreatePayInvoice(ctx, "Нотальная карта", "Оплата услуги", 8500)
		return ctx.Send(invoice)
	})
	bot.Handle("/test", func(ctx tele.Context) error {
		invoice := payment.CreatePayInvoice(ctx, "Прогноз по звездам", "Оплата услуги", 7000)
		return ctx.Send(invoice)
	})

	// bot.Handle(tele.OnText, func(ctx tele.Context) error {
	// 	switch text := ctx.Text(); text {
	// 	case "Нотальная карта":
	// 		err := payment.SendPayInvoice(ctx, "Нотальная карта", "Оплата услуги", 8500)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 			return nil
	// 		}
	// 	case "Еще какая то рандомная хрень":
	// 		err := payment.SendPayInvoice(ctx, "Еще какая то рандомная хрень", "Оплата услуги", 10000)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 			return nil
	// 		}
	// 	default:
	// 		return ctx.Send("Введите одну из предложенных комманд или нажмите на кнопку")
	// 	}
	// 	return ctx.Send("У вас есть 10 минут до конца действительности чека")
	// })
	bot.Start()
}
