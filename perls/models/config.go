package models

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


type Config struct {
	Database DatabaseConfig
}
