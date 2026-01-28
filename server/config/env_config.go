package config

import (
	"os"
)

type appConfig struct {
	DatabaseURL         string
	AccessSecret        string
	RefreshSecret       string
	CloudinaryCloudName string
	CloudinaryApiKey    string
	CloudinaryApiSecret string
}

func AppConfig() appConfig {
	DATABASE_URL := os.Getenv("DATABASE_URL")
	ACCESS_SECRET := os.Getenv("ACCESS_SECRET")
	REFRESH_SECRET := os.Getenv("REFRESH_SECRET")
	CLOUDINARY_CLOUD_NAME := os.Getenv("CLOUDINARY_CLOUD_NAME")
	CLOUDINARY_API_KEY := os.Getenv("CLOUDINARY_API_KEY")
	CLOUDINARY_API_SECRET := os.Getenv("CLOUDINARY_API_SECRET")

	return appConfig{
		DatabaseURL:         DATABASE_URL,
		AccessSecret:        ACCESS_SECRET,
		RefreshSecret:       REFRESH_SECRET,
		CloudinaryCloudName: CLOUDINARY_CLOUD_NAME,
		CloudinaryApiKey:    CLOUDINARY_API_KEY,
		CloudinaryApiSecret: CLOUDINARY_API_SECRET,
	}
}
