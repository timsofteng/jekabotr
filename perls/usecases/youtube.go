package usecases

import (
	"jekabot/models"
	"log"
	"math/rand"
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

var orderValues = [6]string{"date", "rating", "relevance", "title", "videoCount", "viewCount"}

func (u *myYtUsecases) GetRandomVideoUrl() (url string, err error) {
	var fns []func(n int) string

	fns = append(fns, RandEnStringRunes)
	fns = append(fns, RandUaStringRunes)

	randQueryFunc := fns[rand.Intn(len(fns))]

	randomIndex := rand.Intn(len(orderValues))
	order := orderValues[randomIndex]

	retries := 4

	var id string

	for retries > 0 {
		randQuery := randQueryFunc(3)
		id, err = u.repo.GetVideoUrl(randQuery, order)
		log.Println(id)
		if id == "" {
			log.Println("cannot find video")
			retries -= 1
		} else if err != nil {
			return
		} else {
			break
		}
	}

	log.Println(id)

	if len(id) < 1 {
		u.GetRandomVideoUrl()
		return
	}

	url = "https://www.youtube.com/watch?v=" + id

	return
}
