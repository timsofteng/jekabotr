package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

var db *pgx.Conn
var ctx context.Context

func initDB() {
	var err error
	ctx = context.Background()

	db, err = pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	// defer db.Close(context.Background())
}

func getRandMessage() string {
	var randMessage string

	queryRandMessage := `SELECT text FROM text_messages ORDER BY RANDOM() LIMIT 1;`

	err := db.QueryRow(ctx, queryRandMessage).Scan(&randMessage)

	log.Print(randMessage)

	if err != nil {
		log.Printf("QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return randMessage
}

func getMessagesCount() int {
	var count int

	queryCount := `SELECT count(*) FROM text_messages`

	err := db.QueryRow(ctx, queryCount).Scan(&count)

	log.Print("messages in db: ", count)

	if err != nil {
		log.Printf("QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return count
}
