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

func dbGetRandMessage() string {
	var randMessage string

	queryRandMessage := `SELECT text FROM messages ORDER BY RANDOM() LIMIT 1;`

	err := db.QueryRow(ctx, queryRandMessage).Scan(&randMessage)

	log.Print(randMessage)

	if err != nil {
		log.Printf("QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return randMessage
}

func dbGetMessagesCount() int {
	var count int

	queryCount := `SELECT count(*) FROM messages`

	err := db.QueryRow(ctx, queryCount).Scan(&count)

	log.Print("messages in db: ", count)

	if err != nil {
		log.Printf("QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return count
}

func dbAddMessage(message string) {
	query := "INSERT INTO messages (text) VALUES ($1)"

	_, err := db.Exec(ctx, query, message)

	if err != nil {
		log.Printf("Adding failed: %v\n", err)
		os.Exit(1)
	}

	log.Print("message added to database: ", message)
}
