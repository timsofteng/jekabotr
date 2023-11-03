package models

type VoiceMessage struct {
	ID        int64  `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type VoiceMessageRepository interface {
	GetRandVoiceMessage() (string, error)
	GetVoiceMessagesCount() (int, error)
	AddVoiceId(voiceId string) error
}

type VoiceMessageUsecases interface {
	GetRandVoiceMessage() (string, error)
	GetVoiceMessagesCount() (int, error)
	AddVoiceId(voiceId string) error
}
