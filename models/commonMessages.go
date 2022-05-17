package models

type CommonMessagesUsecases interface {
	GetMessagesCount() (int, int, error)
}
