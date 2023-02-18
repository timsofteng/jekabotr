package text

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type TextMessage struct {
	ID        int64  `json:"id"`
	Data      string `json:"data"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type TextMessageRepository interface {
	GetRandTextMessage() (string, error)
	GetTextMessagesCount() (int, error)
	AddTextMessage(message string) error
}

type TextMessageUsecases interface {
	GetRandTextMessage() (string, error)
	AddTextMessage(message string) error
	GetTextMessagesCount() (int, error)
}

type TextDelivery struct {
	TextUs TextMessageUsecases
	Bot    *tgbotapi.BotAPI
}
