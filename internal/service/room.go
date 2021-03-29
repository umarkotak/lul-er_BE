package service

import (
	"context"
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/umarkotak/lul-er_BE/internal/models"
	"github.com/umarkotak/lul-er_BE/internal/repository"
	"google.golang.org/appengine/log"
)

func CreateRoom(reqCreateRoom models.Room) (models.Room, error) {

}