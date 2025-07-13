package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"bot/components/chatgpt"
	"bot/components/keyboards"
	message "bot/components/message"
	"bot/components/payment"
	r "bot/components/redis"
	configReader "bot/config"

	tele "gopkg.in/telebot.v4"
)

func main() {
	bot_token := configReader.Readconfig().BOTTOKEN
	pref := tele.Settings{
		Token:  bot_token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	//?* СОЗДАНИЕ РЕДИС КЛИЕНТА
	RedisClient := r.NewClient()
	redisCtx := context.Background()

	if err := RedisClient.Setter(redisCtx, "State", "default", 10*time.Minute); err != nil {
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
		currentState, err := RedisClient.Getter(redisCtx, "State")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		key := strconv.Itoa(int(ctx.Sender().ID))
		userData, err := RedisClient.GetUser(redisCtx, key)
		if err != nil {
			panic(err)
		}
		ctx.Send("Отправляю запрос, ожидайте...")
		switch currentState {
		case "starPay":
			text := fmt.Sprintf(message.StarRequest, userData.Info)
			resp := chatgpt.RequestOpenAi(text)
			maxLen := 4096
			for i := 0; i < len(resp); i += maxLen {
				end := i + maxLen
				if end > len(resp) {
					end = len(resp)
				}
				part := resp[i:end]
				ctx.Send(part)
			}
			return ctx.Send("Обращайтесь еще!")
		case "notalPay":
			text := fmt.Sprintf(message.NotalMap, userData.Info)
			resp := chatgpt.RequestOpenAi(text)
			maxLen := 4096
			for i := 0; i < len(resp); i += maxLen {
				end := i + maxLen
				if end > len(resp) {
					end = len(resp)
				}
				part := resp[i:end]
				ctx.Send(part)
			}
			return ctx.Send("Обращайтесь еще!")
		default:
			user := ctx.Sender()
			text := fmt.Sprintf("%s, Оплата прошла успешно", user.Username)
			return ctx.Send(text)
		}
	})

	bot.Handle("/start", func(ctx tele.Context) error {

		if err := RedisClient.SetNewUser(redisCtx, int(ctx.Sender().ID), ctx.Sender().Username, false, " ", nil, 24*30*time.Hour); err != nil {
			panic(err)
		}

		if err := RedisClient.Setter(redisCtx, "State", "infoWait", 5*time.Minute); err != nil {
			panic(err)
		}
		return ctx.Send(message.StartText)
	})
	//?* Обработчик доп информации о пользователе
	bot.Handle(tele.OnText, func(ctx tele.Context) error {
		state, err := RedisClient.Getter(redisCtx, "State")
		if err != nil {
			panic(err)
		}
		if state == "infoWait" {
			var (
				userInfo = ctx.Text()
			)
			if err := RedisClient.UpdateFieldUser(redisCtx, int(ctx.Sender().ID), "info", userInfo, 24*30*time.Hour); err != nil {
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
		if err := RedisClient.Setter(redisCtx, "State", "starPay", 10*time.Minute); err != nil {
			panic(err)
		}
		invoice := payment.CreatePayInvoice(ctx, "Прогноз по звездам", "Оплата услуги", 7000)
		return ctx.Send(invoice)
	})
	bot.Handle(&keyboards.BtnAnotherBuy, func(ctx tele.Context) error {
		if err := RedisClient.Setter(redisCtx, "State", "notalPay", 10*time.Minute); err != nil {
			panic(err)
		}
		invoice := payment.CreatePayInvoice(ctx, "Еще какая то штука", "оплата услуги", 10000)
		return ctx.Send(invoice)
	})
	bot.Handle(&keyboards.BtnNotalCard, func(ctx tele.Context) error {
		if err := RedisClient.Setter(redisCtx, "State", "notalPay", 10*time.Minute); err != nil {
			panic(err)
		}
		invoice := payment.CreatePayInvoice(ctx, "Нотальная карта", "Оплата услуги", 8500)
		return ctx.Send(invoice)
	})
	bot.Handle("/test", func(ctx tele.Context) error {
		if err := RedisClient.Setter(redisCtx, "State", "starPay", 10*time.Minute); err != nil {
			panic(err)
		}
		invoice := payment.CreatePayInvoice(ctx, "Прогноз по звездам", "Оплата услуги", 0)
		return ctx.Send(invoice)
	})
	bot.Handle("/testPay", func(ctx tele.Context) error {

		key := strconv.Itoa(int(ctx.Sender().ID))
		userInformation, err := RedisClient.Getter(redisCtx, key)
		if err != nil {
			panic(err)
		}
		text := fmt.Sprintf(message.TaroAdvice, userInformation)
		resp := chatgpt.RequestOpenAi(text)
		return ctx.Send(resp, &tele.SendOptions{
			ParseMode: "Markdown",
		})
	})
	bot.Start()
}
