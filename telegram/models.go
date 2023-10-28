package main

type Config struct {
	Token       string
	BotSign     string
	JekaRealid  string
	PavelychRealId  string
	SouceChatID string
}


type CombinedConfig struct {
	Dev Config
	Prod Config
}
