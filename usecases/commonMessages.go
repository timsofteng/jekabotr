package usecases

import (
	"jekabot/models"
)

type myCommonMessagesUsecases struct {
	TextDB  models.TextMessageRepository
	VoiceDB models.VoiceMessageRepository
}

func NewCommonMessagesUsecases(
	textDB models.TextMessageRepository, voiceDB models.VoiceMessageRepository) models.CommonMessagesUsecases {
	return &myCommonMessagesUsecases{
		TextDB:  textDB,
		VoiceDB: voiceDB,
	}
}

func (u *myCommonMessagesUsecases) GetMessagesCount() (text int, voice int, err error) {
	text, err = u.TextDB.GetTextMessagesCount()
	voice, err = u.VoiceDB.GetVoiceMessagesCount()

	if err != nil {
		return
	}

	return
}
