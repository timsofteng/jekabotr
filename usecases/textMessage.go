package usecases

import (
	"jekabot/models"
)

type myTextUsecases struct {
	DB models.TextMessageRepository
}

func NewTextUsecases(
	db models.TextMessageRepository) models.TextMessageUsecases {
	return &myTextUsecases{
		DB: db,
	}
}

func (u *myTextUsecases) GetRandTextMessage() (randMsg string, err error) {
	randMsg, err = u.DB.GetRandTextMessage()
	if err != nil {
		return
	}

	return
}

func (u *myTextUsecases) AddTextMessage(message string) (err error) {
	err = u.DB.AddTextMessage(message)
	if err != nil {
		return
	}

	return
}
