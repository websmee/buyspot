package telegram

import (
	"log"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func TestAPI(t *testing.T) {
	bot, err := tgbotapi.NewBotAPI("5796928035:AAHFeSu_bt405Ql5DOkglj_WjYsH2T9Spf4")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//updates := bot.GetUpdatesChan(u)
	//
	//for update := range updates {
	//	log.Println(update.Message.From.UserName, update.Message.Chat.ID)
	//}
	bot.Send(tgbotapi.NewMessage(-1001963992438, "test"))
}
