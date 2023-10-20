package config

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var opt = option.WithCredentialsFile("config/serviceAccountKey.json")

func runInit() *firebase.App {
	app, err := firebase.NewApp(context.Background(), &firebase.Config{ProjectID: "artics-3755a"}, opt)
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

func Storage() *storage.BucketHandle {

	storage, err := firebaseApp.Storage(context.Background())
	if err != nil {
		log.Fatalf("error initializing storage app: %v\n", err)
	}
	bucket, err := storage.Bucket("artics-3755a.appspot.com")
	if err != nil {
		log.Fatalf("error initializing storage bucket: %v\n", err)
	}

	return bucket
}
