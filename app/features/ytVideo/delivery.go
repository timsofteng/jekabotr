package ytVideo

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const YT_LINK_CAPTION = "Взгляните на это видео:\n\n"

func initDelivery(
	ytUs YoutubeUsecases,
	bot *tgbotapi.BotAPI) *YoutubeDelivery {

	return &YoutubeDelivery{
		YtUs: ytUs,
		Bot:  bot,
	}

}

func (d *YoutubeDelivery) RespondWithYtUrl(update tgbotapi.Update) {
	ytUrl, err := d.YtUs.GetRandomVideoUrl()
	if err != nil {
		log.Printf("yt url error: %v", err)
	}

	msgText := YT_LINK_CAPTION + ytUrl

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
	msg.ReplyToMessageID = update.Message.MessageID
	d.Bot.Send(msg)
}
