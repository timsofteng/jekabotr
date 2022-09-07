package models

type TelegramConfig struct {
	Token       string
	BotSign     string
	JekaRealid  string
	PavelychRealId  string
	SouceChatID string
}

type DatabaseConfig struct {
	Type                 string
	User                 string
	Password             string
	Net                  string
	Addr                 string
	DBName               string
	AllowNativePasswords bool
	Params               struct {
		ParseTime string
	}
}

type ImgApiConfig struct {
	BaseUrl  string
	ClientId string
}

type YoutubeApiConfig struct {
	Key string
}

type Config struct {
	Database DatabaseConfig
	Telegram TelegramConfig
	ImgApi ImgApiConfig
	YoutubeApi YoutubeApiConfig
}
