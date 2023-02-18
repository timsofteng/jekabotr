package ytVideo

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type YoutubeRepository interface {
	GetVideoUrl(query string, order string) (string, error)
}

type YoutubeUsecases interface {
	GetRandomVideoUrl() (string, error)
}

type YoutubeDelivery struct {
	YtUs     YoutubeUsecases
	Bot      *tgbotapi.BotAPI
}
