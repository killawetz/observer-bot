package telegram

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"sort"
	"strings"
)

const (
	commandMyStat = "mystat"
	topStat       = "topstat"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {

	switch message.Command() {

	case commandMyStat:
		return b.handleMyStatCommand(message)
	case topStat:
		return b.handleTopStatCommand(message)
	default:
		return b.handleMessage(message)
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	//check if exist
	b.addNew(message)

	if message.Audio != nil {
		sql := "update chat_users set audio=audio+1 where chat_id=$1 and member_id=$2"
		b.repo.DBPool.QueryRow(context.Background(), sql, message.Chat.ID, message.From.ID).Scan()
	}
	if message.Sticker != nil {
		sql := "update chat_users set sticker=sticker+1 where chat_id=$1 and member_id=$2"
		b.repo.DBPool.QueryRow(context.Background(), sql, message.Chat.ID, message.From.ID).Scan()
	}
	if message.Video != nil {
		sql := "update chat_users set video=video+1 where chat_id=$1 and member_id=$2"
		b.repo.DBPool.QueryRow(context.Background(), sql, message.Chat.ID, message.From.ID).Scan()
	}
	if message.VideoNote != nil {
		sql := "update chat_users set video_note=chat_users.video_note+1 where chat_id=$1 and member_id=$2"
		b.repo.DBPool.QueryRow(context.Background(), sql, message.Chat.ID, message.From.ID).Scan()
	}
	if message.Photo != nil {
		sql := "update chat_users set photo=photo+1 where chat_id=$1 and member_id=$2"
		b.repo.DBPool.QueryRow(context.Background(), sql, message.Chat.ID, message.From.ID).Scan()
	}
	if message.Document != nil {
		sql := "update chat_users set doc=doc+1 where chat_id=$1 and member_id=$2"
		b.repo.DBPool.QueryRow(context.Background(), sql, message.Chat.ID, message.From.ID).Scan()
	}
	if message.Animation != nil {
		sql := "update chat_users set animation=animation+1 where chat_id=$1 and member_id=$2"
		b.repo.DBPool.QueryRow(context.Background(), sql, message.Chat.ID, message.From.ID).Scan()
	}
	if message.Voice != nil {
		sql := "update chat_users set voice=voice+1 where chat_id=$1 and member_id=$2"
		b.repo.DBPool.QueryRow(context.Background(), sql, message.Chat.ID, message.From.ID).Scan()
	}
	if message.Text != "" {
		fmt.Println(message.Text)
		sql := "update chat_users set text_message=text_message+1 where chat_id=$1 and member_id=$2"
		b.repo.DBPool.QueryRow(context.Background(), sql, message.Chat.ID, message.From.ID).Scan()
	}
	return nil
}

func (b *Bot) handleMyStatCommand(message *tgbotapi.Message) error {

	query := `SELECT cu.*, u.username, u.firstname, u.lastname
	from chat_users cu inner join users u on cu.member_id = u.id
	where chat_id = $1 and member_id = $2`

	var us []*UserStat
	err := pgxscan.Select(context.Background(), b.repo.DBPool, &us, query, message.Chat.ID, message.From.ID)

	if err != nil {
		fmt.Println("Ошибка при попытке получить chat_users: ", err)
	}

	if len(us) < 1 {
		b.addNew(message)
		err := pgxscan.Select(context.Background(), b.repo.DBPool, &us, query, message.Chat.ID, message.From.ID)
		if err != nil {
			fmt.Println("Ошибка при попытке получить chat_users : ", err)
		}
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, us[0].ToString())
	msg.ReplyToMessageID = message.MessageID
	_, err = b.API.Send(msg)
	if err != nil {
		fmt.Println("Ошибка при отправке сообщения: ", err)
	}

	return nil
}

func (b *Bot) handleTopStatCommand(message *tgbotapi.Message) error {
	var us []*UserStat
	query := `SELECT cu.*, u.username, u.firstname, u.lastname
			FROM chat_users cu
			inner join users u on cu.member_id = u.id
			where chat_id = $1`
	err := pgxscan.Select(context.Background(), b.repo.DBPool, &us, query, message.Chat.ID)
	if err != nil {
		fmt.Println("Ошибка при попытке получения статистики всех пользователей чата: ", err)
	}

	sort.Slice(us[:], func(i, j int) bool {
		return us[i].TextMessage > us[j].TextMessage
	})

	var sb strings.Builder
	sb.WriteString("Топ писак:\n")

	for i, el := range us {
		str := fmt.Sprintf("%d. %s (%s): %d сообщений\n", i+1, el.Firstname, el.Username, el.TotalMsgCount())
		sb.WriteString(str)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, sb.String())

	_, err = b.API.Send(msg)
	if err != nil {
		fmt.Println("Ошибка при отправке ответа на /topstat", err)
	}
	return nil
}
