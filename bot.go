package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"bot/components/chatgpt"
	"bot/components/keyboards"
	message "bot/components/message"
	"bot/components/payment"
	r "bot/components/redis"
	configReader "bot/config"

	tele "gopkg.in/telebot.v4"
)

// Local Imports

func main() {
	bot_token := configReader.Readconfig().BOTTOKEN
	pref := tele.Settings{
		Token:  bot_token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	//?* СОЗДАНИЕ РЕДИС КЛИЕНТА
	RedisClient := r.NewClient()

	if err := RedisClient.Setter(context.Background(), "State", "default", 10*time.Minute); err != nil {
		fmt.Println(err)
		os.Exit(1)
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
	// TODO СДЕЛАТЬ ВОЗВРАТ СРЕДСТВ
	bot.Handle("/refund", func(ctx tele.Context) error {
		return ctx.Send("Эта Функция находится в разработке")
	})
	bot.Handle(tele.OnPayment, func(ctx tele.Context) error {
		currentState, err := RedisClient.Getter(context.Background(), "State")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		userInformation, err := RedisClient.Getter(context.Background(), ctx.Sender().Username)
		if err != nil {
			panic(err)
		}
		switch currentState {
		case "starPay":
			text := fmt.Sprintf(message.StarRequest, userInformation)
			resp := chatgpt.RequestOpenAi(text)
			return ctx.Send(resp)
		case "notalPay":
			text := fmt.Sprintf(message.NotalMap, userInformation)
			resp := chatgpt.RequestOpenAi(text)
			return ctx.Send(resp)
		default:
			user := ctx.Sender()
			text := fmt.Sprintf("%s, Оплата прошла успешно", user.Username)
			return ctx.Send(text)
		}
	})
	bot.Handle("/start", func(ctx tele.Context) error {
		if err := RedisClient.Setter(context.Background(), "State", "infoWait", 5*time.Minute); err != nil {
			panic(err)
		}
		return ctx.Send(message.StartText)
	})
	//?* Обработчик доп информации о пользователе
	bot.Handle(tele.OnText, func(ctx tele.Context) error {
		state, err := RedisClient.Getter(context.Background(), "State")
		if err != nil {
			panic(err)
		}
		if state == "infoWait" {
			var (
				userInfo = ctx.Text()
				user     = ctx.Sender().Username
			)
			if err := RedisClient.Setter(context.Background(), user, userInfo, 24*time.Hour); err != nil {
				panic(err)
			}
			return ctx.Send("Спасибо за информацию! Я ее запомню")
		}
		return ctx.Send("Упс, меня не научили говорить :)")
	})
	bot.Handle("/buy", func(ctx tele.Context) error {
		return ctx.Send(message.BuyText, menu)
	})

	bot.Handle(&keyboards.BtnStarCard, func(ctx tele.Context) error {
		if err := RedisClient.Setter(context.Background(), "State", "starPay", 10*time.Minute); err != nil {
			panic(err)
		}
		invoice := payment.CreatePayInvoice(ctx, "Прогноз по звездам", "Оплата услуги", 7000)
		return ctx.Send(invoice)
	})
	bot.Handle(&keyboards.BtnAnotherBuy, func(ctx tele.Context) error {
		if err := RedisClient.Setter(context.Background(), "State", "notalPay", 10*time.Minute); err != nil {
			panic(err)
		}
		invoice := payment.CreatePayInvoice(ctx, "Еще какая то штука", "оплата услуги", 10000)
		return ctx.Send(invoice)
	})
	bot.Handle(&keyboards.BtnNotalCard, func(ctx tele.Context) error {
		if err := RedisClient.Setter(context.Background(), "State", "notalPay", 10*time.Minute); err != nil {
			panic(err)
		}
		invoice := payment.CreatePayInvoice(ctx, "Нотальная карта", "Оплата услуги", 8500)
		return ctx.Send(invoice)
	})
	bot.Handle("/test", func(ctx tele.Context) error {
		if err := RedisClient.Setter(context.Background(), "State", "starPay", 10*time.Minute); err != nil {
			panic(err)
		}
		invoice := payment.CreatePayInvoice(ctx, "Прогноз по звездам", "Оплата услуги", 1)
		return ctx.Send(invoice)
	})
	bot.Start()
}
