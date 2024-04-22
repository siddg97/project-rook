package services

import (
	"context"
	"fmt"
	"github.com/siddg97/project-rook/models"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/rs/zerolog/log"
	"google.golang.org/api/option"
)

type FirebaseService struct {
	ctx             context.Context
	FirebaseApp     *firebase.App
	FirestoreClient *firestore.Client
}

var resumeCollection = "resumes"
var promptHistorySubCollection = "promptHistory"

func InitializeFirebase(ctx context.Context, firebaseConfigPath string) (*FirebaseService, error) {
	opt := option.WithCredentialsFile(firebaseConfigPath)

	firebaseApp, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Error().Msgf("Error creating Firebase app: %v", err)
		return nil, err
	}

	client, err := firebaseApp.Firestore(ctx)
	if err != nil {
		log.Error().Msgf("Error creating Firestore client: %v", err)
		return nil, err
	}

	firebaseService := &FirebaseService{
		ctx:             ctx,
		FirestoreClient: client,
		FirebaseApp:     firebaseApp,
	}

	log.Info().Msg("Successfully initialized Firestore client")
	return firebaseService, nil
}

func (s *FirebaseService) StoreNewResume(userId string, resumeInText string) error {
	resumeDocRef := s.FirestoreClient.Collection(resumeCollection).Doc(userId)
	_, err := resumeDocRef.Set(s.ctx, &models.ResumeDocument{
		UserID:     userId,
		ResumeID:   userId,
		ResumeText: resumeInText,
	})
	if err != nil {
		log.Err(err).Msg("Failed to write resume document to firebase")
		return err
	}

	promptHistoryDocRef := resumeDocRef.Collection(promptHistorySubCollection).NewDoc()
	_, err = promptHistoryDocRef.Set(s.ctx, &models.PromptHistoryDocument{
		Id:        promptHistoryDocRef.ID,
		CreatedAt: time.Now(),
		Role:      "rook",
		Text:      fmt.Sprintf("Providing the content of the resume as text:\n\n%s\n\nKeep this in mind as you will be provided with updates of what the person has done since this resume and asked to update relevant sections of the resume", resumeInText),
	})
	if err != nil {
		log.Err(err).Msgf("Failed to write to prompt history for user: %s", userId)
		return err
	}
	return nil
}

func (s *FirebaseService) GetResume(userId string) (*models.ResumeDocument, error) {
	resumeDocRef, err := s.FirestoreClient.Collection(resumeCollection).Doc(userId).Get(s.ctx)
	if err != nil {
		return nil, err
	}
	var resumeDocument models.ResumeDocument
	resumeDocRef.DataTo(&resumeDocument)
	return &resumeDocument, nil
}
