package services

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/api/option"
	"google.golang.org/api/vision/v1"
)

var visionService vision.Service

func InitializeVision(ctx context.Context) (*vision.Service, error) {
	service, err := vision.NewService(ctx, option.WithScopes(vision.CloudPlatformScope))
	if err != nil {
		log.Fatal().Msgf("Error creating Vision Service Client: %v", err)
		return nil, err
	}

	visionService = *service

	log.Info().Msg("Successfully initialized Vision Service client")

	return service, nil
}

func GetVisionService() vision.Service {
	return visionService
}
