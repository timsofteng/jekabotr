package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"strconv"
	"strings"
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

	if trigger {
		// randomMessageNumber := rand.Intn(linesCount)
		// text := ReadExactLine(randomMessageNumber)

		randMsg := dbGetRandMessage()

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, randMsg)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

func tgRun() {
	// linesCount := countNumbers()
	// rand.Seed(time.Now().UnixNano())

	// log.Println("messages in bot:", linesCount)

	tgInit()
	tgWatchUpdates()

	for update := range updates {
		if update.Message != nil { // If we got a message
			chatId := update.FromChat().ID
			strCtatId := strconv.Itoa(int(chatId))

			if strCtatId == sourceChatId {
				dbAddMessage(update.Message.Text)
			} else {
				tgRespond(update)
			}
		}
	}
}
