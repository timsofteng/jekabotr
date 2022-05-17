package models

type Taksa struct {
	Urls struct{ Full string }
	Id   string
}

type TaksaRepository interface {
	GetRandomTaksaUrl() (string, string, error)
	GetBytesFromUrl(url string) ([]byte, error)
}

type TaksaUsecases interface {
	GetRandomTaksa() ([]byte, string, error)
}
