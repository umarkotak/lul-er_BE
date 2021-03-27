package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/umarkotak/lul-er_BE/internal/app"
)

func main() {
	fmt.Println("Welcome to lul-er tangga backend service")

	initEnvConfiguration()

	app.Start()
}

func initEnvConfiguration() {
	godotenv.Load(".env")
}
