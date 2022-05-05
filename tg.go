package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var token string
var botSign string
var sourceChatId string
var bot *tgbotapi.BotAPI
var updates tgbotapi.UpdatesChannel

var err error

func tgInit() {
	token = os.Getenv("TG_TOKEN")
	botSign = os.Getenv("BOT_SIGN")
	sourceChatId = os.Getenv("SOURCE_CHAT_ID")

	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
}

func tgWatchUpdates() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates = bot.GetUpdatesChan(u)
}

func tgRespond(update tgbotapi.Update) {
	textMsg := update.Message.Text
	author := update.Message.From.UserName
	isReply := update.Message.ReplyToMessage
	var isReplyToBot bool

	if isReply != nil {
		replyTo := update.Message.ReplyToMessage.From.UserName
		isReplyToBot = replyTo == botSign
	}
	log.Printf("[%s] %s \n", author, textMsg)

	isTriggerWords := strings.Contains(strings.ToLower(textMsg), "jeka")
	isAuthorJeka := author == "Jekadesigner"
	trigger := isTriggerWords || isAuthorJeka || isReplyToBot

	//make rundomize for text messages more
	//get rid of this piece of shit
	var fns []func(update tgbotapi.Update)
	fns = append(fns, tgRespondRandText)
	fns = append(fns, tgRespondRandText)
	fns = append(fns, tgRespondRandText)
	fns = append(fns, tgRespondRandText)
	fns = append(fns, tgRespondRandText)
	fns = append(fns, tgRespondRandVoice)

	randFunc := fns[rand.Intn(len(fns))]

	if trigger {
		randFunc(update)
	}
}

func tgRespondRandText(update tgbotapi.Update) {
	randMsg := dbGetRandTextMessage()
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, randMsg)
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}

func tgRespondRandVoice(update tgbotapi.Update) {
	voiceId := dbGetRandVoiceMessage()
	voice := tgbotapi.NewVoice(update.Message.Chat.ID, tgbotapi.FileID(voiceId))
	voice.ReplyToMessageID = update.Message.MessageID
	bot.Send(voice)
}

func tgRun() {
	// linesCount := countNumbers()
	// rand.Seed(time.Now().UnixNano())

	// log.Println("messages in bot:", linesCount)

	tgInit()
	tgWatchUpdates()
	rand.Seed(time.Now().UTC().UnixNano())

	for update := range updates {
		if update.Message != nil { // If we got a message
			// log.Print(update.Message.Voice.FileUniqueID)
			chatId := update.FromChat().ID
			strChattId := strconv.Itoa(int(chatId))

			if strChattId == sourceChatId {
				if update.Message.Voice != nil {
					voiceId := update.Message.Voice.FileID
					dbAddVoiceId(voiceId)
				} else {
					dbAddTextMessage(update.Message.Text)
				}
			} else {
				tgRespond(update)
			}
		}
	}
}
