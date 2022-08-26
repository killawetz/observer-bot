package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/killawetz/observer-bot/pkg/config"
	"github.com/killawetz/observer-bot/pkg/storage"
	"github.com/killawetz/observer-bot/pkg/telegram"
	"log"
)

// init is invoked before main()
func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	conf := config.New()

	botAPI, err := tgbotapi.NewBotAPI(conf.TelegramAPIKey)
	if err != nil {
		log.Panic(err)
	}

	repo := storage.New(conf.ConnString)

	bot := telegram.NewBot(botAPI, repo)

	bot.Start()

}
