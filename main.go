package main

import (
	"fmt"
	"jekabot/config"
	"jekabot/delivery"
	repo "jekabot/repository"
	us "jekabot/usecases"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	c := config.ReadConfig()
	cDB := c.Database
	cTg := c.Telegram

	dbConn := fmt.Sprintf("%s://%s@%s/%s", cDB.Type, cDB.User, cDB.Addr, cDB.DBName)
	db := repo.NewDB(dbConn)

	textRepo := repo.NewTextRepository(db)
	voiceRepo := repo.NewVoiceRepository(db)
	taksaRepo := repo.NewTaksaRepository(c.ImgApi.BaseUrl, c.ImgApi.ClientId)

	textUs := us.NewTextUsecases(textRepo)
	voiceUs := us.NewVoiceUsecases(voiceRepo)
	taksaUs := us.NewTaksaUsecases(taksaRepo)
	commonUs := us.NewCommonMessagesUsecases(textRepo, voiceRepo)

	bot, err := tgbotapi.NewBotAPI(cTg.Token)
	if err != nil {
		log.Panic(err)
	}

	d := delivery.NewDelivery(
		textUs,
		voiceUs,
		commonUs,
		taksaUs,
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
