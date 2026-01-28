package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/Sahil2004/gmarket/server/config"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type ImageService struct {
	cld *cloudinary.Cloudinary
	ctx context.Context
}

func NewImageService() (*ImageService, error) {
	cld, err := cloudinary.NewFromParams(config.AppConfig().CloudinaryCloudName, config.AppConfig().CloudinaryApiKey, config.AppConfig().CloudinaryApiSecret)
	if err != nil {
		return nil, fmt.Errorf("cloudinary init failed: %w", err)
	}

	cld.Config.URL.Secure = true

	ctx := context.Background()

	return &ImageService{
		cld: cld,
		ctx: ctx,
	}, nil
}

func (s *ImageService) UploadFromDataURL(dataURL string, publicID string) (*uploader.UploadResult, error) {
	if !strings.HasPrefix(dataURL, "data:") {
		return nil, fmt.Errorf("invalid data url")
	}

	resp, err := s.cld.Upload.Upload(s.ctx, dataURL, uploader.UploadParams{
		PublicID:       publicID,
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true),
		ResourceType:   "image",
	})

	if err != nil {
		return nil, fmt.Errorf("upload from data url failed: %w", err)
	}

	return resp, nil
}
