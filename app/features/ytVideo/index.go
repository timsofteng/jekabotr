package ytVideo

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func GetTaksaDelivery(ytApiApiKey string, bot *tgbotapi.BotAPI) *YoutubeDelivery {
	taksaRepo := NewYoutubeRepository(ytApiApiKey)
	taksaUs := NewYoutubeUsecases(taksaRepo)

	d := initDelivery(taksaUs, bot)
	return d
}
