package usecases

import (
	"log"
	"youtube/models"
	"youtube/randQuery"
)

type myYtUsecases struct {
	repo models.YoutubeRepository
}

const YT_LINK_CAPTION = "Взгляните на это видео:\n\n"

func NewYoutubeUsecases(
	repo models.YoutubeRepository) models.YoutubeUsecases {
	return &myYtUsecases{
		repo: repo,
	}
}

func (u *myYtUsecases) GetRandomVideoUrl() (id string, caption string, err error) {
	retries := 4
	caption = YT_LINK_CAPTION

	for retries > 0 {
		randQueryStr := randQuery.String(3)
		randOrder := randQuery.Order()
		id, err = u.repo.GetVideoUrl(randQueryStr, randOrder)

		if err != nil {
			return id, caption, err
		}

		if id == "" {
			log.Printf("cannot find video, will try again %v times", retries)
			retries -= 1
			continue
		}

		break
	}

	url := "https://www.youtube.com/watch?v=" + id

	return url, caption, nil
}
