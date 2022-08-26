package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UserStat struct {
	ChatId      int    `db:"chat_id"`
	MemberId    int    `db:"member_id"`
	TextMessage int    `db:"text_message"`
	Sticker     int    `db:"sticker"`
	Voice       int    `db:"voice"`
	Audio       int    `db:"audio"`
	Video       int    `db:"video"`
	VideoNote   int    `db:"video_note"`
	Doc         int    `db:"doc"`
	Animation   int    `db:"animation"`
	Photo       int    `db:"photo"`
	Username    string `db:"username"`
	Firstname   string `db:"firstname"`
	Lastname    string `db:"lastname"`
}

func (us *UserStat) ToString() string {
	str := fmt.Sprintf("Стастика пользователя %s %s (@%s)\n\n"+
		"Кол-во сообщений: 	%d\n"+
		"Кол-во стикеров: 	%d\n"+
		"Кол-во голосовых: 	%d\n"+
		"Кол-во аудио: 		%d\n"+
		"Кол-во видео: 		%d\n"+
		"Кол-во кружочков: 	%d\n"+
		"Кол-во доков: 		%d\n"+
		"Кол-во гифок: 		%d\n"+
		"Кол-во фото:		%d\n",
		us.Firstname, us.Lastname, us.Username,
		us.TotalMsgCount(), us.Sticker, us.Voice, us.Audio, us.Video, us.VideoNote, us.Doc, us.Animation, us.Photo)
	return str
}

func (us *UserStat) TotalMsgCount() int {
	return us.TextMessage + us.Sticker + us.Voice + us.Audio + us.VideoNote + us.Video + us.Doc + us.Animation
}

func (b *Bot) addNew(message *tgbotapi.Message) {
	chatID := message.Chat.ID
	userID := message.From.ID

	queryToChat := `select exists(select * from chat where id=$1)`
	queryToUsers := `select exists(select * from users where id=$1)`

	var isExist bool
	err := b.repo.DBPool.QueryRow(context.Background(), queryToChat, chatID).Scan(&isExist)
	if err != nil {
		fmt.Println("Something went wrong when trying to access the CHAT table")
	}

	if !isExist {
		insertQuery := `insert into chat values ($1, $2)`
		// check the success of the insertion?
		b.repo.DBPool.QueryRow(context.Background(), insertQuery, chatID, message.Chat.Title)
	}

	err = b.repo.DBPool.QueryRow(context.Background(), queryToUsers, userID).Scan(&isExist)
	if err != nil {
		fmt.Println("Something went wrong when trying to access the USERS table ")
	}

	if !isExist {
		insertQuery := `insert into users values ($1, $2, $3, $4)`
		// check the success of the insertion?
		b.repo.DBPool.QueryRow(context.Background(),
			insertQuery,
			userID,
			message.From.UserName,
			message.From.FirstName,
			message.From.LastName).Scan()
	}

	queryToChatUsers := `select exists(select * from chat_users where chat_id=$2 and member_id=$1)`

	err = b.repo.DBPool.QueryRow(context.Background(), queryToChatUsers, chatID, userID).Scan(&isExist)
	if err != nil {
		fmt.Println("Something went wrong when trying to access the USERS table ")
	}

	if !isExist {
		insertQuery := `insert into chat_users values ($1, $2)`
		// check the success of the insertion?
		b.repo.DBPool.QueryRow(context.Background(), insertQuery, chatID, userID).Scan()
	}

}
