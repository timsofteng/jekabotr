package usecases

import (
	"log"
	models "youtube/models"
)

type myYtUsecases struct {
	repo models.YoutubeRepository
}

func NewYoutubeUsecases(
	repo models.YoutubeRepository) models.YoutubeUsecases {
	return &myYtUsecases{
		repo: repo,
	}
}

func (u *myYtUsecases) GetRandomVideoUrl() (string, error) {
	retries := 4

	var id string
	var err error

	for retries > 0 {
		randQuery := RandString(3)
		randOrder := RandOrder()
		id, err = u.repo.GetVideoUrl(randQuery, randOrder)

		if err != nil {
			return id, err
		}

		if id == "" {
			log.Printf("cannot find video, will try again %v times", retries)
			retries -= 1
			continue
		}

		break
	}

	url := "https://www.youtube.com/watch?v=" + id

	return url, nil
}
