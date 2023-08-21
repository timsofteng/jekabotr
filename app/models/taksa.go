package models

type TaksaRepository interface {
	GetRandomTaksaUrl() (string, string, error)
}

type TaksaUsecases interface {
	GetRandomTaksa() ([]byte, string, error)
}
