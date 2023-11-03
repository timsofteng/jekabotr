package usecases

import (
	models "images/models"
)

type myTaksaUsecases struct {
	repo models.ImagesRepository
}

const TAKSA_CAPTION = "Собака умная может и самоутилизироваться )\n😍😍😍😍"

func NewTaksaUsecases(
	repo models.ImagesRepository) models.ImagesUsecases {
	return &myTaksaUsecases{
		repo: repo,
	}
}

func (u myTaksaUsecases) RandomTaksaGetter() ([]byte, string, string, error) {
	bin, id, err := u.repo.ImgByQueryFetcher("dachshund")

	if err != nil {
		return nil, id, "", err
	}

	return bin, id, TAKSA_CAPTION, err
}
