package models

type DbMethods interface {
	GetRandTextMessage() (string, error)
	GetRandVoiceMessage() (string, error)
	GetTextMessagesCount() (int, error)
	GetVoiceMessagesCount() (int, error)
	AddTextMessage(message string) error
	AddVoiceId(voiceId string) error
}
