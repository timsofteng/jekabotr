package repository

import (
	"log"
	"os"

	"jekabot/models"

	"github.com/jackc/pgx/v4/pgxpool"
)


type myVoiceRepo struct {
	conn *pgxpool.Pool
}

func NewVoiceRepository(db *pgxpool.Pool) models.VoiceMessageRepository {
	return &myVoiceRepo{conn: db}
	// defer db.Close(context.Background())
}


func (r *myVoiceRepo) GetRandVoiceMessage() (randVoiceId string, err error) {
	query := `SELECT id FROM voice ORDER BY RANDOM() LIMIT 1;`

	err = r.conn.QueryRow(ctx, query).Scan(&randVoiceId)

	if err != nil {
		return
	}

	return
}

func (r *myVoiceRepo) GetVoiceMessagesCount() (count int, err error) {
	query := `SELECT count(*) FROM voice`

	err = r.conn.QueryRow(ctx, query).Scan(&count)

	if err != nil {
		log.Printf("QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return
}

func (r *myVoiceRepo) AddVoiceId(voiceId string) (err error) {
	query := "INSERT INTO voice (id) VALUES ($1)"

	_, err = r.conn.Exec(ctx, query, voiceId)

	if err != nil {
		log.Printf("Adding failed: %v\n", err)
		os.Exit(1)
	}

	log.Print("voice_id added to database: ", voiceId)

	return
}

