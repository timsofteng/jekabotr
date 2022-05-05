package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

var db *pgx.Conn
var ctx context.Context

func dbInit() {
	var err error
	ctx = context.Background()

	db, err = pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	// defer db.Close(context.Background())
}

func dbGetRandTextMessage() string {
	var randMessage string

	query := `SELECT data FROM text ORDER BY RANDOM() LIMIT 1;`

	err := db.QueryRow(ctx, query).Scan(&randMessage)

	log.Print(randMessage)

	if err != nil {
		log.Printf("QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return randMessage
}

func dbGetRandVoiceMessage() string {
	var randVoiceId string

	query := `SELECT id FROM voice ORDER BY RANDOM() LIMIT 1;`

	err := db.QueryRow(ctx, query).Scan(&randVoiceId)

	log.Print(randVoiceId)

	if err != nil {
		log.Printf("QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return randVoiceId
}

func dbGetTextMessagesCount() int {
	var count int

	query := `SELECT count(*) FROM text`

	err := db.QueryRow(ctx, query).Scan(&count)

	log.Print("messages in db: ", count)

	if err != nil {
		log.Printf("QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return count
}

func dbAddTextMessage(message string) {
	query := "INSERT INTO text (data) VALUES ($1)"

	_, err := db.Exec(ctx, query, message)

	if err != nil {
		log.Printf("Adding failed: %v\n", err)
		os.Exit(1)
	}

	log.Print("message added to database: ", message)
}

func dbAddVoiceId(voiceId string) {
	query := "INSERT INTO voice (id) VALUES ($1)"

	_, err := db.Exec(ctx, query, voiceId)

	if err != nil {
		log.Printf("Adding failed: %v\n", err)
		os.Exit(1)
	}

	log.Print("voice_id added to database: ", voiceId)
}
