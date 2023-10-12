package config

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/storage"
	"google.golang.org/api/option"
)

var opt = option.WithCredentialsFile("/serviceAccountKey.json")

func runInit() *firebase.App {
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v\n", err)
	}
	return app
}

var firebaseApp = runInit()

func Firestore() *firestore.Client {

	firestore, err := firebaseApp.Firestore(context.Background())
	if err != nil {
		log.Fatalf("error initializing firestore app: %v\n", err)
	}
	return firestore

}

func Auth() *auth.Client {

	auth, err := firebaseApp.Auth(context.Background())
	if err != nil {
		log.Fatalf("error initializing auth app: %v\n", err)
	}
	return auth
}

func Storage() *storage.Client {

	storage, err := firebaseApp.Storage(context.Background())
	if err != nil {
		log.Fatalf("error initializing storage app: %v\n", err)
	}
	return storage
}
