package config

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/storage"
)

func runInit() *firebase.App {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v\n", err)
	}
	return app
}

func Firestore() *firestore.Client {
	firebaseApp := runInit()
	firestore, err := firebaseApp.Firestore(context.Background())
	if err != nil {
		log.Fatalf("error initializing firestore app: %v\n", err)
	}
	return firestore

}

func Auth() *auth.Client {
	firebaseApp := runInit()
	auth, err := firebaseApp.Auth(context.Background())
	if err != nil {
		log.Fatalf("error initializing auth app: %v\n", err)
	}
	return auth
}

func Storage() *storage.Client {
	firebaseApp := runInit()
	storage, err := firebaseApp.Storage(context.Background())
	if err != nil {
		log.Fatalf("error initializing storage app: %v\n", err)
	}
	return storage
}
