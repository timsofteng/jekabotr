package models

type ImagesRepository interface {
	ImgByQueryFetcher(query string) ([]byte, string, error)
}

type ImagesUsecases interface {
	RandomTaksaGetter() (bin []byte, id string, caption string, err error)
}
