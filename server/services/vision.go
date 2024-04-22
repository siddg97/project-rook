package services

import (
	"context"
	"encoding/base64"
	"github.com/rs/zerolog/log"
	"google.golang.org/api/option"
	"google.golang.org/api/vision/v1"
)

type VisionService struct {
	visionClient vision.Service
	ctx          context.Context
}

func InitializeVision(ctx context.Context) (*VisionService, error) {
	service, err := vision.NewService(ctx, option.WithScopes(vision.CloudPlatformScope))
	if err != nil {
		log.Err(err).Msgf("Error creating Vision Service Client: %v", err)
		return nil, err
	}

	visionService := &VisionService{
		visionClient: *service,
		ctx:          ctx,
	}

	log.Info().Msg("Successfully initialized Vision Service client")

	return visionService, nil
}

func (s *VisionService) ExtractTextFromPdf(pdfBytes []byte) (string, error) {
	request := &vision.BatchAnnotateFilesRequest{
		Requests: []*vision.AnnotateFileRequest{
			{
				InputConfig: &vision.InputConfig{
					MimeType: "application/pdf",
					Content:  base64.StdEncoding.EncodeToString(pdfBytes),
				},
				Features: []*vision.Feature{
					{Type: "DOCUMENT_TEXT_DETECTION"},
				},
			},
		},
	}
	response, err := s.visionClient.Files.Annotate(request).Context(s.ctx).Do()
	if err != nil {
		log.Err(err).Msgf("Error when calling vision client to annotate PDF files")
		return "", err
	}

	// Process GCP vision output
	var extractedTextFromPdf string
	for _, page := range response.Responses[0].Responses {
		extractedTextFromPdf += page.FullTextAnnotation.Text
	}

	log.Info().Msg("Successfully extracted text from PDF file")
	return extractedTextFromPdf, nil
}
