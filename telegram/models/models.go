package models

type Config struct {
	Token       string
	BotSign     string
	JekaRealid  string
	PavelychRealId  string
	SouceChatID string
}


type MultiEnvConfig struct {
	Dev Config
	Prod Config
}
