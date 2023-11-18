package delivery

import (
	"context"
	"log"
	"strconv"
	"strings"
	"telegram/models"

	pb "github.com/jeka-designer/proto/gen/go"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const errMsgToUser = "шо-то пошло по пизде... давай сначала"

type myDelivery struct {
	ytGRPCClient     pb.YoutubeServiceClient
	imagesGRPCClient pb.ImagesServiceClient
	config           *models.Config
	bot              *tgbotapi.BotAPI
	Updates          tgbotapi.UpdatesChannel
}

func NewDelivery(
	ytGRPCClient pb.YoutubeServiceClient,
	imagesGRPCClient pb.ImagesServiceClient,
	c *models.Config,
) (*myDelivery, error) {

	bot, err := tgbotapi.NewBotAPI(c.Token)
	if err != nil {
		return nil, err
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// textMsgs, err := textUs.GetTextMessagesCount()
	// voiceMsgs, err := voiceUs.GetVoiceMessagesCount()
	// textMsgsStr := strconv.Itoa(int(textMsgs))
	// voiceMsgsStr := strconv.Itoa(int(voiceMsgs))

	if err != nil {
		log.Printf("total count messages err: %v", err)
	}

	// log.Printf("total text messages: %s   total voices: %s", textMsgsStr, voiceMsgsStr)
	// bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	return &myDelivery{
		// textUs:  textUs,
		// voiceUs: voiceUs,
		ytGRPCClient:     ytGRPCClient,
		imagesGRPCClient: imagesGRPCClient,
		config:           c,
		bot:              bot,
		Updates:          updates,
	}, err

}

func (d *myDelivery) Router(update tgbotapi.Update) error {
	chatId := update.FromChat().ID
	strChattId := strconv.Itoa(int(chatId))

	if strChattId != d.config.SouceChatID {
		err := d.respRouter(update)
		return err
	} else {
		d.storeRouter(update)
	}

	return nil

}

func (t *myDelivery) respRouter(update tgbotapi.Update) error {

	textMsg := update.Message.Text
	author := update.Message.From.UserName

	log.Printf("[%s] %s \n", author, textMsg)

	if strings.Contains(strings.ToLower(textMsg), "jeka_taksa") {
		err := t.RespondWithTaksa(update)
		return err
	}

	if strings.Contains(strings.ToLower(textMsg), "jeka_video") {
		err := t.RespondWithYtUrl(update)
		return err
	}

	// isReply := update.Message.ReplyToMessage

	// var isReplyToBot bool

	// if isReply != nil {
	// 	replyTo := update.Message.ReplyToMessage.From.UserName
	// 	isReplyToBot = replyTo == t.config.BotSign
	// }

	// isTriggerWords := strings.Contains(strings.ToLower(textMsg), "jeka")
	// isAuthorJeka := author == t.config.JekaRealid
	// isAuthorPavelych := author == t.config.PavelychRealId
	// trigger := isTriggerWords || isAuthorJeka || isAuthorPavelych || isReplyToBot

	//make rundomize for text messages properly
	//get rid of this piece of shit
	// var fns []func(update tgbotapi.Update)
	// fns = append(fns, t.RespondWithText)
	// fns = append(fns, t.RespondWithText)
	// fns = append(fns, t.RespondWithText)
	// fns = append(fns, t.RespondWithText)
	// fns = append(fns, t.RespondWithText)
	// fns = append(fns, t.RespondWithText)
	// fns = append(fns, t.RespondWithVoice)

	// randFunc := fns[rand.Intn(len(fns))]

	// if trigger {
	// 	randFunc(update)
	// }
	return nil
}

func (d *myDelivery) storeRouter(update tgbotapi.Update) {
	if update.Message.Voice != nil {
		// voiceId := update.Message.Voice.FileID
		// d.voiceUs.AddVoiceId(voiceId)
	} else {
		// d.textUs.AddTextMessage(update.Message.Text)
	}
}

func (d *myDelivery) respondWithError(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, errMsgToUser)
	_, _ = d.bot.Send(msg)
}

func (d *myDelivery) RespondWithTaksa(update tgbotapi.Update) error {
	log.Println("Random taksa called")

	resp, err := d.imagesGRPCClient.GetRandomTaksa(context.Background(), &pb.GetRandomTaksaRequest{})
	if err != nil {
		d.respondWithError(update)
		return err
	}

	msg := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileBytes{Name: resp.Id, Bytes: resp.Bin})
	msg.ReplyToMessageID = update.Message.MessageID
	msg.Caption = resp.Caption

	_, err = d.bot.Send(msg)

	if err != nil {
		d.respondWithError(update)
		return err
	}

	log.Print("responded with some image")

	return nil
}

func (d *myDelivery) RespondWithYtUrl(update tgbotapi.Update) error {
	log.Println("Random video called")

	resp, err := d.ytGRPCClient.GetRandomVideo(context.Background(), &pb.GetRandomVideoRequest{})
	if err != nil {
		d.respondWithError(update)
		return err
	}

	msgText := resp.Caption + resp.Url

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
	msg.ReplyToMessageID = update.Message.MessageID

	_, err = d.bot.Send(msg)

	if err != nil {
		d.respondWithError(update)
		return err
	}

	return nil
}

// func (d *myDelivery) RespondWithText(update tgbotapi.Update) {
// 	randMsg, err := d.textUs.GetRandTextMessage()
// 	if err != nil {
// 		log.Printf("rand text error: %v", err)
// 	}

// 	msg := tgbotapi.NewMessage(update.Message.Chat.ID, randMsg)
// 	msg.ReplyToMessageID = update.Message.MessageID
// 	d.bot.Send(msg)
// }

// func (d *myDelivery) RespondWithVoice(update tgbotapi.Update) {
// 	voiceId, err := d.voiceUs.GetRandVoiceMessage()
// 	if err != nil {
// 		log.Printf("rand voice error: %v", err)
// 	}
// 	voice := tgbotapi.NewVoice(update.Message.Chat.ID, tgbotapi.FileID(voiceId))
// 	voice.ReplyToMessageID = update.Message.MessageID
// 	d.bot.Send(voice)
// }
