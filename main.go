package main

import (
	"fmt"
	"github.com/killawetz/observer-bot/config"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

// init is invoked before main()
func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	conf := config.New()

	bot, err := tgbotapi.NewBotAPI(conf.APIKey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			/*msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID
			*/
			inputMessage := update.Message

			msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "")

			switch inputMessage.Command() {
			case "my_stat":
				msg.Text = fmt.Sprintf("Статистика участника %s: %d", inputMessage.From.UserName, 0)
			}
			bot.Send(msg)
		}
	}
}
