package main

import (
	"fmt"
	"jekabot/delivery"
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

	dbConn := fmt.Sprintf("%s://%s:%s@%s/%s",
		cDB.Type,
		cDB.User,
		cDB.Password,
		cDB.Addr,
		cDB.DBName)

	log.Println(dbConn)

	db := repo.NewDB(dbConn)

	textRepo := repo.NewTextRepository(db)
	voiceRepo := repo.NewVoiceRepository(db)
	taksaRepo := repo.NewTaksaRepository(c.ImgApi.BaseUrl, c.ImgApi.ClientId)
	ytRepo := repo.NewYoutubeRepository(cYt.Key)

	textUs := us.NewTextUsecases(textRepo)
	voiceUs := us.NewVoiceUsecases(voiceRepo)
	taksaUs := us.NewTaksaUsecases(taksaRepo)
	ytUs := us.NewYoutubeUsecases(ytRepo)

	bot, err := tgbotapi.NewBotAPI(cTg.Token)
	if err != nil {
		log.Panic(err)
	}

	d := delivery.NewDelivery(
		textUs,
		voiceUs,
		taksaUs,
		ytUs,
		cTg,
		bot,
	)

	// bot.Debug = true

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
