package services

import (
	"context"
	"sort"
	"time"

	"github.com/siddg97/project-rook/models"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/rs/zerolog/log"
)

type FirebaseService struct {
	ctx             context.Context
	FirebaseApp     *firebase.App
	FirestoreClient *firestore.Client
}

const resumeCollection = "resumes"
const promptHistorySubCollection = "promptHistory"

func InitializeFirebase(ctx context.Context) (*FirebaseService, error) {
	firebaseApp, err := firebase.NewApp(ctx, nil)
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

	return nil
}

func (s *FirebaseService) GetResumePromptHistory(userId string) ([]models.PromptHistoryDocument, error) {
	promptHistorySubCollectionRef := s.FirestoreClient.Collection(resumeCollection).Doc(userId).Collection(promptHistorySubCollection)
	promptHistory, err := promptHistorySubCollectionRef.Documents(s.ctx).GetAll()
	if err != nil {
		log.Err(err).Msgf("Error retrieving prompt history for user %s", userId)
		return make([]models.PromptHistoryDocument, 0), err
	}

	var promptHistoryDocs []models.PromptHistoryDocument
	for _, docRef := range promptHistory {
		var doc models.PromptHistoryDocument
		docRef.DataTo(&doc)
		promptHistoryDocs = append(promptHistoryDocs, doc)
	}

	sort.Slice(promptHistoryDocs, func(i, j int) bool {
		return promptHistoryDocs[i].CreatedAt.Before(promptHistoryDocs[j].CreatedAt)
	})

	return promptHistoryDocs, nil
}

func (s *FirebaseService) StoreToPromptHistory(userId string, promptText string, role string, input string) error {
	promptHistorySubCollectionRef := s.FirestoreClient.Collection(resumeCollection).Doc(userId).Collection(promptHistorySubCollection)

	latestPromptDocRef := promptHistorySubCollectionRef.NewDoc()
	_, err := latestPromptDocRef.Set(s.ctx, &models.PromptHistoryDocument{
		Id:        latestPromptDocRef.ID,
		CreatedAt: time.Now(),
		Role:      role,
		Text:      promptText,
		Input:     input,
	})
	if err != nil {
		log.Err(err).Msgf("Failed to write to prompt history for user: %s", userId)
		return err
	}

	log.Info().Msgf("Saved prompt for role: %s to prompt history: %s", role, promptText)

	return nil
}
func (s *FirebaseService) GetResume(userId string) (*models.ResumeDocument, error) {
	resumeDocRef, err := s.FirestoreClient.Collection(resumeCollection).Doc(userId).Get(s.ctx)
	if err != nil {
		return nil, err
	}
	var resumeDocument models.ResumeDocument
	err = resumeDocRef.DataTo(&resumeDocument)
	if err != nil {
		return nil, err
	}
	return &resumeDocument, nil
}
