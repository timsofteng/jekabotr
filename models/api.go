package models

type RandomSingleImg struct {
	Urls struct{ Full string }
	Id   string
}

type ApiMethods interface {
	GetRandomTaksaUrl() (string, string, error)
	GetBytesFromUrl(url string) ([]byte, error)
}
