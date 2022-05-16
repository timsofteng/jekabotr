package models

type Usecases interface {
	GetRandTextMessage() (string, error)
	GetRandVoiceMessage() (string, error)
	GetMessagesCount() (int, int, error)
	AddTextMessage(message string) error
	AddVoiceId(voiceId string) error
	GetRandomTaksa() ([]byte, string, error)
}

