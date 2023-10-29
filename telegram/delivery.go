package main

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"strings"

	pb "proto"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type myDelivery struct {
	ytClient pb.YoutubeServiceClient
	config   *Config
	bot      *tgbotapi.BotAPI
	Updates  tgbotapi.UpdatesChannel
}

const TAKSA_CAPTION = "Собака умная может и самоутилизироваться )\n😍😍😍😍"
const YT_LINK_CAPTION = "Взгляните на это видео:\n\n"

func NewDelivery(
	ytClient pb.YoutubeServiceClient,
	c *Config,
) *myDelivery {

	bot, err := tgbotapi.NewBotAPI(c.Token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// textMsgs, err := textUs.GetTextMessagesCount()
	// voiceMsgs, err := voiceUs.GetVoiceMessagesCount()
	// textMsgsStr := strconv.Itoa(int(textMsgs))
	// voiceMsgsStr := strconv.Itoa(int(voiceMsgs))

	if err != nil {
		log.Printf("total count messages err: %v", err)
	}

	// log.Printf("total text messages: %s   total voices: %s", textMsgsStr, voiceMsgsStr)
	// bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	return &myDelivery{
		// textUs:  textUs,
		// voiceUs: voiceUs,
		// taksaUs: taksaUs,
		ytClient: ytClient,
		config:   c,
		bot:      bot,
		Updates:  updates,
	}

}

func (d *myDelivery) Router(update tgbotapi.Update) {
	chatId := update.FromChat().ID
	strChattId := strconv.Itoa(int(chatId))

	if strChattId != d.config.SouceChatID {
		d.respRouter(update)
	} else {
		d.storeRouter(update)
	}

}

func (t *myDelivery) respRouter(update tgbotapi.Update) {

	textMsg := update.Message.Text
	author := update.Message.From.UserName

	log.Printf("[%s] %s \n", author, textMsg)

	// if strings.Contains(strings.ToLower(textMsg), "jeka_taksa") {
	// 	t.RespondWithTaksa(update)
	// 	return
	// }

	if strings.Contains(strings.ToLower(textMsg), "jeka_video") {
		t.RespondWithYtUrl(update)
		return
	}

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
	// fns = append(fns, t.RespondWithText)
	// fns = append(fns, t.RespondWithText)
	// fns = append(fns, t.RespondWithText)
	// fns = append(fns, t.RespondWithText)
	// fns = append(fns, t.RespondWithText)
	// fns = append(fns, t.RespondWithText)
	// fns = append(fns, t.RespondWithVoice)

	randFunc := fns[rand.Intn(len(fns))]

	if trigger {
		randFunc(update)
	}
}

func (d *myDelivery) storeRouter(update tgbotapi.Update) {
	if update.Message.Voice != nil {
		// voiceId := update.Message.Voice.FileID
		// d.voiceUs.AddVoiceId(voiceId)
	} else {
		// d.textUs.AddTextMessage(update.Message.Text)
	}
}

// func (d *myDelivery) RespondWithTaksa(update tgbotapi.Update) {

// 	bytes, id, err := d.taksaUs.GetRandomTaksa()
// 	if err != nil {
// 		log.Printf("rand taksa error: %v", err)
// 	}

// 	msg := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileBytes{Name: id, Bytes: bytes})
// 	msg.ReplyToMessageID = update.Message.MessageID
// 	msg.Caption = TAKSA_CAPTION
// 	d.bot.Send(msg)
// }

func (d *myDelivery) RespondWithYtUrl(update tgbotapi.Update) {
	resp, err := d.ytClient.GetRandomVideo(context.Background(), &pb.GetRandomVideoRequest{})
	if err != nil {
		log.Printf("yt url error: %v", err)
	}

	msgText := YT_LINK_CAPTION + resp.Url

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
	msg.ReplyToMessageID = update.Message.MessageID
	d.bot.Send(msg)
}

// func (d *myDelivery) RespondWithText(update tgbotapi.Update) {
// 	randMsg, err := d.textUs.GetRandTextMessage()
// 	if err != nil {
// 		log.Printf("rand text error: %v", err)
// 	}

// 	msg := tgbotapi.NewMessage(update.Message.Chat.ID, randMsg)
// 	msg.ReplyToMessageID = update.Message.MessageID
// 	d.bot.Send(msg)
// }

// func (d *myDelivery) RespondWithVoice(update tgbotapi.Update) {
// 	voiceId, err := d.voiceUs.GetRandVoiceMessage()
// 	if err != nil {
// 		log.Printf("rand voice error: %v", err)
// 	}
// 	voice := tgbotapi.NewVoice(update.Message.Chat.ID, tgbotapi.FileID(voiceId))
// 	voice.ReplyToMessageID = update.Message.MessageID
// 	d.bot.Send(voice)
// }
