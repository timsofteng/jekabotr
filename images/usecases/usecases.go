package usecases
import (
	models "images/models"
)

type myTaksaUsecases struct {
	repo models.ImagesRepository
}

func NewTaksaUsecases(
	repo models.ImagesRepository) models.ImagesUsecases {
	return &myTaksaUsecases{
		repo: repo,
	}
}

func (u *myTaksaUsecases) GetRandomTaksa() ([]byte, string, error) {
	url, id, err := u.repo.GetImgByQuery("dachshund")

	if err != nil {
		return nil, id, err
	}

	bin, err := BytesFromUrl(url)

	if err != nil {
		return bin, id, err
	}

	return bin, id, err
}
