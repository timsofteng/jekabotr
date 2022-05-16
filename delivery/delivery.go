package delivery

import (
	"jekabot/models"
	"log"
	"math/rand"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Delivery struct {
	Usecases models.Usecases
	Bot      *tgbotapi.BotAPI
	Config   models.TelegramConfig
}

const TAKSA_CAPTION = "Собака умная может и самоутилизироваться )\n😍😍😍😍"

func NewDelivery(c models.TelegramConfig, usecases models.Usecases, bot *tgbotapi.BotAPI) *Delivery {

	return &Delivery{
		Usecases: usecases,
		Bot:      bot,
		Config:   c,
	}

}

func (t *Delivery) Router(update tgbotapi.Update) {
	chatId := update.FromChat().ID
	strChattId := strconv.Itoa(int(chatId))

	if strChattId != t.Config.SouceChatID {
		t.respRouter(update)
	} else {
		t.storeRouter(update)
	}

}

func (t *Delivery) respRouter(update tgbotapi.Update) {

	textMsg := update.Message.Text
	author := update.Message.From.UserName

	log.Printf("[%s] %s \n", author, textMsg)

	if strings.Contains(strings.ToLower(textMsg), "jeka_taksa") {
		go t.RespondWithTaksa(update)
		return
	}
	isReply := update.Message.ReplyToMessage

	var isReplyToBot bool

	if isReply != nil {
		replyTo := update.Message.ReplyToMessage.From.UserName
		isReplyToBot = replyTo == t.Config.BotSign
	}

	isTriggerWords := strings.Contains(strings.ToLower(textMsg), "jeka")
	isAuthorJeka := author == t.Config.JekaRealid
	trigger := isTriggerWords || isAuthorJeka || isReplyToBot

	//make rundomize for text messages more
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

func (t *Delivery) storeRouter(update tgbotapi.Update) {
	if update.Message.Voice != nil {
		voiceId := update.Message.Voice.FileID
		t.Usecases.AddVoiceId(voiceId)
	} else {
		t.Usecases.AddTextMessage(update.Message.Text)
	}
}

func (t *Delivery) RespondWithTaksa(update tgbotapi.Update) {

	bytes, id, err := t.Usecases.GetRandomTaksa()
	if err != nil {
	}

	msg := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileBytes{Name: id, Bytes: bytes})
	msg.ReplyToMessageID = update.Message.MessageID
	msg.Caption = TAKSA_CAPTION
	t.Bot.Send(msg)
}

func (t *Delivery) RespondWithText(update tgbotapi.Update) {
	randMsg, err := t.Usecases.GetRandTextMessage()
	if err != nil {
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, randMsg)
	msg.ReplyToMessageID = update.Message.MessageID
	t.Bot.Send(msg)
}

func (t *Delivery) RespondWithVoice(update tgbotapi.Update) {
	voiceId, err := t.Usecases.GetRandVoiceMessage()
	if err != nil {
	}
	voice := tgbotapi.NewVoice(update.Message.Chat.ID, tgbotapi.FileID(voiceId))
	voice.ReplyToMessageID = update.Message.MessageID
	t.Bot.Send(voice)
}
