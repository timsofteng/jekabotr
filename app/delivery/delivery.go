package delivery

import (
	taksa "jekabot/features/taksa"
	text "jekabot/features/text"
	ytVideo "jekabot/features/ytVideo"
	"jekabot/models"
	"log"
	"math/rand"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MyDelivery struct {
	TextDelivery  *text.TextDelivery
	VoiceUs       models.VoiceMessageUsecases
	TaksaDelivery *taksa.TaksaDelivery
	YtDelivery    *ytVideo.YoutubeDelivery
	TgConfig      models.TelegramConfig
	Bot           *tgbotapi.BotAPI
}

const TAKSA_CAPTION = "Собака умная может и самоутилизироваться )\n😍😍😍😍"
const YT_LINK_CAPTION = "Взгляните на это видео:\n\n"

func NewDelivery(
	args MyDelivery) *MyDelivery {

	textMsgs, err := args.TextDelivery.GetTextMessagesCount()
	voiceMsgs, err := args.VoiceUs.GetVoiceMessagesCount()

	textMsgsStr := strconv.Itoa(int(textMsgs))
	voiceMsgsStr := strconv.Itoa(int(voiceMsgs))

	if err != nil {
		log.Printf("total count messages err: %v", err)
	}

	log.Printf("total text messages: %s   total voices: %s", textMsgsStr, voiceMsgsStr)

	return &MyDelivery{
		TextDelivery:  args.TextDelivery,
		VoiceUs:       args.VoiceUs,
		TaksaDelivery: args.TaksaDelivery,
		YtDelivery:    args.YtDelivery,
		TgConfig:      args.TgConfig,
		Bot:           args.Bot,
	}

}

func (d *MyDelivery) Router(update tgbotapi.Update) {
	chatId := update.FromChat().ID
	strChattId := strconv.Itoa(int(chatId))

	if strChattId != d.TgConfig.SouceChatID {
		d.respRouter(update)
	} else {
		d.storeRouter(update)
	}

}

func (t *MyDelivery) respRouter(update tgbotapi.Update) {

	textMsg := update.Message.Text
	author := update.Message.From.UserName

	log.Printf("[%s] %s \n", author, textMsg)

	if strings.Contains(strings.ToLower(textMsg), "jeka_taksa") {
		t.TaksaDelivery.RespondWithTaksa(update)
		return
	}

	if strings.Contains(strings.ToLower(textMsg), "jeka_video") {
		t.YtDelivery.RespondWithYtUrl(update)
		return
	}

	isReply := update.Message.ReplyToMessage

	var isReplyToBot bool

	if isReply != nil {
		replyTo := update.Message.ReplyToMessage.From.UserName
		isReplyToBot = replyTo == t.TgConfig.BotSign
	}

	isTriggerWords := strings.Contains(strings.ToLower(textMsg), "jeka")
	isAuthorJeka := author == t.TgConfig.JekaRealid
	isAuthorPavelych := author == t.TgConfig.PavelychRealId
	trigger := isTriggerWords || isAuthorJeka || isAuthorPavelych || isReplyToBot

	//make rundomize for text messages more
	//get rid of this piece of shit
	var fns []func(update tgbotapi.Update)
	fns = append(fns, t.TextDelivery.RespondWithText)
	fns = append(fns, t.TextDelivery.RespondWithText)
	fns = append(fns, t.TextDelivery.RespondWithText)
	fns = append(fns, t.TextDelivery.RespondWithText)
	fns = append(fns, t.TextDelivery.RespondWithText)
	fns = append(fns, t.TextDelivery.RespondWithText)
	fns = append(fns, t.RespondWithVoice)

	randFunc := fns[rand.Intn(len(fns))]

	if trigger {
		randFunc(update)
	}
}

func (d *MyDelivery) storeRouter(update tgbotapi.Update) {
	if update.Message.Voice != nil {
		voiceId := update.Message.Voice.FileID
		d.VoiceUs.AddVoiceId(voiceId)
	} else {
		d.TextUs.AddTextMessage(update.Message.Text)
	}
}

func (d *MyDelivery) RespondWithVoice(update tgbotapi.Update) {
	voiceId, err := d.VoiceUs.GetRandVoiceMessage()
	if err != nil {
		log.Printf("rand voice error: %v", err)
	}
	voice := tgbotapi.NewVoice(update.Message.Chat.ID, tgbotapi.FileID(voiceId))
	voice.ReplyToMessageID = update.Message.MessageID
	d.Bot.Send(voice)
}
