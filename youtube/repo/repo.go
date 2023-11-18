package repo

import (
	"context"

	"google.golang.org/api/option"
	youtube "google.golang.org/api/youtube/v3"
	"youtube/models"
)

const kievCoordinates = "50.4501, 30.5234"
const maxRadius = "999km"

type myYoutubeRepo struct {
	apiKey string
}

func NewYoutubeRepository(apiKey string) models.YoutubeRepository {
	return &myYoutubeRepo{
		apiKey: apiKey,
	}

}

func (y *myYoutubeRepo) GetVideoUrl(query string, order string) (id string, err error) {

	ctx := context.Background()

	yt, err := youtube.NewService(ctx, option.WithAPIKey(y.apiKey))

	if err != nil {
		return "", err
	}

	var theArray []string
	theArray = append(theArray, "snippet")

	request := yt.Search.List(theArray)

	request.Q(query)

	request.MaxResults(1)

	request.Location(kievCoordinates)
	request.LocationRadius(maxRadius)

	request.Order(order)

	request.Type("video")

	request.RegionCode("ua")

	resp, err := request.Do()

	if err != nil {
		return "", err
	}

	if len(resp.Items) > 0 {
		id = resp.Items[0].Id.VideoId
		return id, err
	}

	return id, err
}
