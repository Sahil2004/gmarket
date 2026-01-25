package config

import (
	"os"
)

type appConfig struct {
	DatabaseURL string
	AccessSecret string
	RefreshSecret string
}

func AppConfig() appConfig {
	DATABASE_URL := os.Getenv("DATABASE_URL")
	ACCESS_SECRET := os.Getenv("ACCESS_SECRET")
	REFRESH_SECRET := os.Getenv("REFRESH_SECRET")

	return appConfig{
		DatabaseURL: DATABASE_URL,
		AccessSecret: ACCESS_SECRET,
		RefreshSecret: REFRESH_SECRET,
	}
}