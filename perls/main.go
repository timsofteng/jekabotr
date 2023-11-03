package main

import (
	"fmt"
	"log"
	"math/rand"
	"perls/delivery"
	repo "perls/repository"
	us "perls/usecases"
	"time"
)

func main() {
	// rand.Seed(time.Now().UnixNano())
	c := ReadConfig()
	cDB := c.Database

	dbConn := fmt.Sprintf("%s://%s:%s@%s/%s",
		cDB.Type,
		cDB.User,
		cDB.Password,
		cDB.Addr,
		cDB.DBName)

	log.Println(dbConn)

	db := repo.NewDB(dbConn)
	defer db.Close()

	textRepo := repo.NewTextRepository(db)
	voiceRepo := repo.NewVoiceRepository(db)

	textUs := us.NewTextUsecases(textRepo)
	voiceUs := us.NewVoiceUsecases(voiceRepo)

	delivery.NewDelivery(
		textUs,
		voiceUs,
	)

}
