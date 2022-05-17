package usecases

import (
	"jekabot/models"
)

type myVoiceMessageUsecases struct {
	DB models.VoiceMessageRepository
}

func NewVoiceUsecases(
	db models.VoiceMessageRepository) models.VoiceMessageUsecases {
	return &myVoiceMessageUsecases{
		DB: db,
	}
}

func (t *myVoiceMessageUsecases) GetRandVoiceMessage() (voiceId string, err error) {
	voiceId, err = t.DB.GetRandVoiceMessage()
	if err != nil {
		return
	}

	return
}

func (t *myVoiceMessageUsecases) AddVoiceId(voiceId string) (err error) {
	err = t.DB.AddVoiceId(voiceId)
	if err != nil {
		return
	}

	return
}
