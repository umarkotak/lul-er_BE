package service

import (
	"context"
	"errors"

	"github.com/umarkotak/lul-er_BE/internal/models"
	"github.com/umarkotak/lul-er_BE/internal/repository"
	"google.golang.org/appengine/log"
)

func Register(reqUser models.User) (models.User, error) {
	user, err := repository.GetUserByUsername(reqUser.Username)
	if err != nil {
		log.Errorf(context.Background(), "Error GetUserByUsername %v", err)
		return user, err
	}

	if user.Username != "" {
		return user, errors.New("username already taken")
	}

	newUser := models.User{
		Username:         reqUser.Username,
		Password:         reqUser.Password,
		Email:            reqUser.Email,
		AuthToken:        "",
		ActiveGameRoomID: "",
	}

	_, err = repository.CreateUser(newUser)
	if err != nil {
		log.Errorf(context.Background(), "Error CreateUser %v", err)
		return user, err
	}

	return newUser, nil
}
