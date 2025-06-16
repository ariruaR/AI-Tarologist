package main

import (
	"context"
	"fmt"
	"os"

	"bot/components/chatgpt"
	"bot/components/keyboards"
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
			bot.SendMessage(ctx,
				ti.Message(
					ti.ID(chatId),
					"Проверяю Звезды...",
				),
			)
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
}
