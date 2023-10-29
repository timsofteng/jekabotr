package models

type ImagesRepository interface {
	GetImgByQuery(query string) (string, string, error)
}

type ImagesUsecases interface {
	GetRandomTaksa() (bin []byte, id string, err error)
}
