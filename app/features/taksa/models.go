package taksa

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Taksa struct {
	Urls struct{ Full string }
	Id   string
}

type TaksaRepository interface {
	GetRandomTaksaUrl() (string, string, error)
	GetBytesFromUrl(url string) ([]byte, error)
}

type TaksaUsecases interface {
	GetRandomTaksa() ([]byte, string, error)
}

type TaksaDelivery struct {
	TaksaUs TaksaUsecases
	Bot     *tgbotapi.BotAPI
}
