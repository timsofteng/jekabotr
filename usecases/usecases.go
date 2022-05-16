package usecases

import (
	"jekabot/models"
)

type UseCases struct {
	db        models.DbMethods
	apiClient models.ApiMethods
}

func NewUsecases(
	db models.DbMethods, apiClient models.ApiMethods) models.Usecases {
	return &UseCases{
		db:        db,
		apiClient: apiClient,
	}
}

func (t *UseCases) GetRandTextMessage() (randMsg string, err error) {
	randMsg, err = t.db.GetRandTextMessage()
	if err != nil {
		return
	}

	return
}

func (t *UseCases) GetRandVoiceMessage() (voiceId string, err error) {
	voiceId, err = t.db.GetRandVoiceMessage()
	if err != nil {
		return
	}

	return
}

func (t *UseCases) GetMessagesCount() (text int, voice int, err error) {
	text, err = t.db.GetTextMessagesCount()
	voice, err = t.db.GetVoiceMessagesCount()

	if err != nil {
		return
	}

	return

}

func (t *UseCases) AddTextMessage(message string) (err error) {
	err = t.db.AddTextMessage(message)
	if err != nil {
		return
	}

	return
}

func (t *UseCases) AddVoiceId(voiceId string) (err error) {
	err = t.db.AddTextMessage(voiceId)
	if err != nil {
		return
	}

	return
}

func (t *UseCases) GetRandomTaksa() (bytes []byte, id string, err error) {
	url, id, err := t.apiClient.GetRandomTaksaUrl()
	if err != nil {
		return 
	}

	bytes, err = t.apiClient.GetBytesFromUrl(url)
	if err != nil {
		return 
	}

	return
}

