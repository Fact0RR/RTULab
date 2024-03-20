package fine

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendFineToTelegram(finalString string, token string,fineprice string,chat string,photo []byte) error {

	intchat, err := strconv.Atoi(chat)

	photoFileBytes := tgbotapi.FileBytes{
		Name:  "fine",
		Bytes: photo,
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}
	photoMsg := tgbotapi.NewPhoto(int64(intchat), photoFileBytes)
	bot.Send(photoMsg)
	msg := tgbotapi.NewMessage(int64(intchat), string(finalString))
	
	bot.Send(msg)
	log.Println("Штраф отправлен")
	return nil
}



