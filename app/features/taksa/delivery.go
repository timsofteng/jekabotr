package taksa

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const TAKSA_CAPTION = "Собака умная может и самоутилизироваться )\n😍😍😍😍"


func initDelivery(
	taksaUs TaksaUsecases,
	bot *tgbotapi.BotAPI) *TaksaDelivery {

	return &TaksaDelivery{
		TaksaUs: taksaUs,
		Bot:     bot,
	}

}

func (d *TaksaDelivery) RespondWithTaksa(update tgbotapi.Update) {

	bytes, id, err := d.TaksaUs.GetRandomTaksa()
	if err != nil {
		log.Printf("rand taksa error: %v", err)
	}

	msg := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileBytes{Name: id, Bytes: bytes})
	msg.ReplyToMessageID = update.Message.MessageID
	msg.Caption = TAKSA_CAPTION
	d.Bot.Send(msg)
}
