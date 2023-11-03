package delivery

import (
	"perls/models"
	"log"
	"math/rand"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type myDelivery struct {
	textUs  models.TextMessageUsecases
	voiceUs models.VoiceMessageUsecases
}

func NewDelivery(
	textUs models.TextMessageUsecases,
	voiceUs models.VoiceMessageUsecases,
) *myDelivery {

	textMsgs, err := textUs.GetTextMessagesCount()
	voiceMsgs, err := voiceUs.GetVoiceMessagesCount()
	textMsgsStr := strconv.Itoa(int(textMsgs))
	voiceMsgsStr := strconv.Itoa(int(voiceMsgs))

	if err != nil {
		log.Printf("total count messages err: %v", err)
	}

	log.Printf("total text messages: %s   total voices: %s", textMsgsStr, voiceMsgsStr)
	// bot.Debug = true

	return &myDelivery{
		textUs:  textUs,
		voiceUs: voiceUs,
	}

}

func (t *myDelivery) respRouter(update tgbotapi.Update) {

	textMsg := update.Message.Text
	author := update.Message.From.UserName

	log.Printf("[%s] %s \n", author, textMsg)


	isReply := update.Message.ReplyToMessage

	var isReplyToBot bool

	if isReply != nil {
		replyTo := update.Message.ReplyToMessage.From.UserName
		isReplyToBot = replyTo == t.config.BotSign
	}

	isTriggerWords := strings.Contains(strings.ToLower(textMsg), "jeka")
	isAuthorJeka := author == t.config.JekaRealid
	isAuthorPavelych := author == t.config.PavelychRealId
	trigger := isTriggerWords || isAuthorJeka || isAuthorPavelych || isReplyToBot

	//make rundomize for text messages properly
	//get rid of this piece of shit
	var fns []func(update tgbotapi.Update)
	fns = append(fns, t.RespondWithText)
	fns = append(fns, t.RespondWithText)
	fns = append(fns, t.RespondWithText)
	fns = append(fns, t.RespondWithText)
	fns = append(fns, t.RespondWithText)
	fns = append(fns, t.RespondWithText)
	fns = append(fns, t.RespondWithVoice)

	randFunc := fns[rand.Intn(len(fns))]

	if trigger {
		randFunc(update)
	}
}

func (d *myDelivery) storeRouter(update tgbotapi.Update) {
	if update.Message.Voice != nil {
		voiceId := update.Message.Voice.FileID
		d.voiceUs.AddVoiceId(voiceId)
	} else {
		d.textUs.AddTextMessage(update.Message.Text)
	}
}

func (d *myDelivery) RespondWithText(update tgbotapi.Update) {
	_, err := d.textUs.GetRandTextMessage()
	if err != nil {
		log.Printf("rand text error: %v", err)
	}

}

func (d *myDelivery) RespondWithVoice(update tgbotapi.Update) {
	_, err := d.voiceUs.GetRandVoiceMessage()
	if err != nil {
		log.Printf("rand voice error: %v", err)
	}
}
