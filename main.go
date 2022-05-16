package main

import (
	"fmt"
	repo "jekabot/repository"
	"jekabot/delivery"
	"jekabot/usecases"
	"jekabot/config"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	c := config.ReadConfig()

	cDB := c.Database

	cTg := c.Telegram

	apiClient := repo.NewClient(c.ImgApi.BaseUrl, c.ImgApi.ClientId)

	dbConn := fmt.Sprintf("%s://%s@%s/%s", cDB.Type, cDB.User, cDB.Addr, cDB.DBName)

	db := repo.NewDB(dbConn)

	bot, err := tgbotapi.NewBotAPI(cTg.Token)
	if err != nil {
		log.Panic(err)
	}

	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	us := usecases.NewUsecases(db, apiClient)
	
	d := delivery.NewDelivery(cTg, us, bot )

	for update := range updates {
		if update.Message != nil { // If we got a message
			d.Router(update)
		}
	}

}
