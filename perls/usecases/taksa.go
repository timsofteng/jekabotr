package usecases

import (
	"jekabot/models"
	"jekabot/utils"
)

type myTaksaUsecases struct {
	repo models.TaksaRepository
}

func NewTaksaUsecases(
	repo models.TaksaRepository) models.TaksaUsecases {
	return &myTaksaUsecases{
		repo: repo,
	}
}

func (u *myTaksaUsecases) GetRandomTaksa() (bytes []byte, id string, err error) {
	url, id, err := u.repo.GetRandomTaksaUrl()
	if err != nil {
		return
	}

	bytes, err = utils.BytesFromUrl(url)
	if err != nil {
		return
	}

	return
}
