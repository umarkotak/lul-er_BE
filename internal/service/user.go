package service

import (
	"context"

	"github.com/umarkotak/lul-er_BE/internal/models"
	"github.com/umarkotak/lul-er_BE/internal/repository"
	"google.golang.org/appengine/log"
)

func Register(reqUser models.User) (models.User, error) {
	user, err := repository.GetUserByUsername(reqUser.Username)
	if err != nil {
		log.Infof(context.Background(), "err %v", err)
	}

	return user, nil
}
