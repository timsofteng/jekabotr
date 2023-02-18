package text

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v4/pgxpool"
)

func GetTextDelivery(db *pgxpool.Pool, bot *tgbotapi.BotAPI) *TextDelivery {
	textRepo := NewTextRepository(db)
	textUs := NewTextUsecases(textRepo)

	d := initDelivery(textUs, bot)
	return d
}

func GetTextUs(db *pgxpool.Pool) TextMessageUsecases {
	textRepo := NewTextRepository(db)
	textUs := NewTextUsecases(textRepo)

	return textUs
}
