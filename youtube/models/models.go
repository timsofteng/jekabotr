package models

type YoutubeRepository interface {
	GetVideoUrl(query string, order string) (string, error)
}

type YoutubeUsecases interface {
	GetRandomVideoUrl() (string, error)
}
