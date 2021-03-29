package config

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

type Config struct {
	FirebaseApp *firebase.App
	FirebaseDB  *db.Client
	FbRootRef   *db.Ref
	FbUsersRef  *db.Ref
}

var config Config

func InitConfig() {
	firebaseApp, firebaseDB := initFirebaseApp()
	fbRootRef := firebaseDB.NewRef("")
	fbUsersRef := fbRootRef.Child("users")

	config = Config{
		FirebaseApp: firebaseApp,
		FirebaseDB:  firebaseDB,
		FbRootRef:   fbRootRef,
		FbUsersRef:  fbUsersRef,
	}
}

func GetConfig() Config {
	return config
}

func initFirebaseApp() (*firebase.App, *db.Client) {
	conf := &firebase.Config{
		DatabaseURL: "https://luler-tangga-default-rtdb.firebaseio.com/",
	}
	opt := option.WithCredentialsFile("internal/config/serviceAccountKey.json")

	firebaseApp, err := firebase.NewApp(context.Background(), conf, opt)
	if err != nil {
		log.Fatalln("Error initializing firebase app:", err)
	}

	firebaseDB, err := firebaseApp.Database(context.Background())
	if err != nil {
		log.Fatalln("Error initializing firebase db:", err)
	}

	return firebaseApp, firebaseDB
}
