package services

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/rs/zerolog/log"
	"google.golang.org/api/option"
)

var firestoreClient *firestore.Client

func InitializeFirestore(ctx context.Context, firebaseConfigPath string) (*firestore.Client, error) {
	opt := option.WithCredentialsFile("firebase_credentials.json")
	firebaseApp, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatal().Msgf("Error creating Firebase app: %v", err)
		return nil, err
	}

	client, err := firebaseApp.Firestore(ctx)
	if err != nil {
		log.Fatal().Msgf("Error creating Firestore client: %v", err)
		return nil, err
	}

	firestoreClient = client

	log.Info().Msg("Successfully initialized Firestore client")
	return firestoreClient, nil
}
