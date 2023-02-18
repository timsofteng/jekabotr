package main

import (
	"jekabot/delivery"
	"jekabot/features/taksa"
	"jekabot/features/text"
	"jekabot/features/ytVideo"
	lib "jekabot/lib"
	repo "jekabot/repository"
	us "jekabot/usecases"
	"log"
	"math/rand"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	c := ReadConfig()
	cDB := c.Database
	cTg := c.Telegram
	cYt := c.YoutubeApi
	cImgApi := c.ImgApi

	db := lib.GetDatabaseInstance(cDB)
	bot := lib.GetTgBotInstance(cTg.Token)

	textDelivery := text.GetTextDelivery(db, bot)
	taksaDelivery := taksa.GetTaksaDelivery(cImgApi.BaseUrl, cImgApi.ClientId, bot)
	ytDelivery := ytVideo.GetTaksaDelivery(cYt.Key, bot)

	textRepo := repo.NewTextRepository(db)
	voiceRepo := repo.NewVoiceRepository(db)

	textUs := us.NewTextUsecases(textRepo)
	voiceUs := us.NewVoiceUsecases(voiceRepo)

	routerArgs := delivery.MyDelivery{
		TextDelivery:  textDelivery,
		VoiceUs:       voiceUs,
		YtDelivery:    ytDelivery,
		TaksaDelivery: taksaDelivery,
		Bot:           bot,
		TgConfig:      cTg,
	}

	d := delivery.NewDelivery(routerArgs)

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			go d.Router(update)
		}
	}

}
