package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/killawetz/observer-bot/pkg/storage"
)

type Bot struct {
	API  *tgbotapi.BotAPI
	repo *storage.Repository
}

func NewBot(bot *tgbotapi.BotAPI, repo *storage.Repository) *Bot {
	return &Bot{
		API:  bot,
		repo: repo,
	}
}

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.API.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		//fmt.Println(update.Message)

		if update.Message.IsCommand() {
			if err := b.handleCommand(update.Message); err != nil {
				b.handleError(update.Message.Chat.ID, err)
			}
			continue
		}

		// Handle regular messages
		if err := b.handleMessage(update.Message); err != nil {
			b.handleError(update.Message.Chat.ID, err)
		}
	}
	return nil
}
