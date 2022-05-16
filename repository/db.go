package repository

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"jekabot/models"
)

var ctx context.Context

type myDb struct {
	conn *pgx.Conn
}

func NewDB(databaseUrl string) models.DbMethods {
	ctx = context.Background()

	db, err := pgx.Connect(ctx, databaseUrl)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return &myDb{conn: db}
	// defer db.Close(context.Background())
}

func (db *myDb) GetRandTextMessage() (randMsg string, err error) {
	query := `SELECT data FROM text ORDER BY RANDOM() LIMIT 1;`

	err = db.conn.QueryRow(ctx, query).Scan(&randMsg)

	log.Print(randMsg)

	if err != nil {
		log.Printf("QueryRow failed: %v\n", err)
		return
	}

	return
}

func (db *myDb) GetRandVoiceMessage() (randVoiceId string, err error) {
	query := `SELECT id FROM voice ORDER BY RANDOM() LIMIT 1;`

	err = db.conn.QueryRow(ctx, query).Scan(&randVoiceId)

	log.Print(randVoiceId)

	if err != nil {
		return
	}

	return
}

func (db *myDb) GetTextMessagesCount() (count int, err error) {
	query := `SELECT count(*) FROM text`

	err = db.conn.QueryRow(ctx, query).Scan(&count)

	if err != nil {
		log.Printf("QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return
}

func (db *myDb) GetVoiceMessagesCount() (count int, err error) {
	query := `SELECT count(*) FROM voice`

	err = db.conn.QueryRow(ctx, query).Scan(&count)

	if err != nil {
		log.Printf("QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return
}

func (db *myDb) AddTextMessage(message string) (err error) {
	query := "INSERT INTO text (data) VALUES ($1)"

	_, err = db.conn.Exec(ctx, query, message)

	if err != nil {
		log.Printf("Adding failed: %v\n", err)
		os.Exit(1)
	}

	log.Print("message added to database: ", message)
	return
}

func (db *myDb) AddVoiceId(voiceId string) (err error) {
	query := "INSERT INTO voice (id) VALUES ($1)"

	_, err = db.conn.Exec(ctx, query, voiceId)

	if err != nil {
		log.Printf("Adding failed: %v\n", err)
		os.Exit(1)
	}

	log.Print("voice_id added to database: ", voiceId)

	return
}
