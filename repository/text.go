package repository

import (
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"jekabot/models"
)


type myTextRepo struct {
	conn *pgx.Conn
}

func NewTextRepository(db *pgx.Conn) models.TextMessageRepository {
	return &myTextRepo{conn: db}
	// defer db.Close(context.Background())
}

func (r *myTextRepo) GetRandTextMessage() (randMsg string, err error) {
	query := `SELECT data FROM text ORDER BY RANDOM() LIMIT 1;`

	err = r.conn.QueryRow(ctx, query).Scan(&randMsg)

	if err != nil {
		return
	}

	return
}

func (r *myTextRepo) GetTextMessagesCount() (count int, err error) {
	query := `SELECT count(*) FROM text`

	err = r.conn.QueryRow(ctx, query).Scan(&count)

	if err != nil {
		log.Printf("QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return
}

func (r *myTextRepo) AddTextMessage(message string) (err error) {
	query := "INSERT INTO text (data) VALUES ($1)"

	_, err = r.conn.Exec(ctx, query, message)

	if err != nil {
		log.Printf("Adding failed: %v\n", err)
		os.Exit(1)
	}

	log.Print("message added to database: ", message)
	return
}
