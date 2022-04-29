package main

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// file, err := os.OpenFile("log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalf("error opening file: %v", err)
	// }
	// defer file.Close()

	// log.SetOutput(file)
	initDB()

	token := os.Getenv("TG_TOKEN")
	botSign := os.Getenv("BOT_SIGN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	linesCount := countNumbers()
	rand.Seed(time.Now().UnixNano())

	log.Println("in main   ", linesCount)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
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
				randomMessageNumber := rand.Intn(linesCount)
				text := ReadExactLine(randomMessageNumber)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			}
		}
	}
}
