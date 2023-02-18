package text

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func initDelivery(
	textUs TextMessageUsecases,
	bot *tgbotapi.BotAPI) *TextDelivery {

	return &TextDelivery{
		TextUs: textUs,
		Bot:    bot,
	}

}

func (d *TextDelivery) RespondWithText(update tgbotapi.Update) {
	randMsg, err := TextUsGetRandTextMessage()
	if err != nil {
		log.Printf("rand text error: %v", err)
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, randMsg)
	msg.ReplyToMessageID = update.Message.MessageID
	d.Bot.Send(msg)
}
