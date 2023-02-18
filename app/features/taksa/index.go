package taksa

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func GetTaksaDelivery(imgApiBaseUrl string, imgApiClienId string, bot *tgbotapi.BotAPI) *TaksaDelivery {
	taksaRepo := NewTaksaRepository(imgApiBaseUrl, imgApiClienId)
	taksaUs := NewTaksaUsecases(taksaRepo)

	d := initDelivery(taksaUs, bot)
	return d
}
