package config

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Config struct {
	firebaseApp *firebase.App
}

var config Config

func InitConfig() {
	config = Config{
		firebaseApp: initFirebaseApp(),
	}
}

func GetConfig() Config {
	return config
}

func initFirebaseApp() *firebase.App {
	firebaseApp, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing firebaseApp: %v\n", err)
	}

	opt := option.WithCredentialsFile("config/serviceAccountKey.json")
	fbConfig := &firebase.Config{ProjectID: "luler-tangga"}
	firebaseApp, err = firebase.NewApp(context.Background(), fbConfig, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return firebaseApp
}
