package repository

import (
	"context"

	"github.com/umarkotak/lul-er_BE/internal/config"
	"github.com/umarkotak/lul-er_BE/internal/models"
)

func GetUserByUsername(username string) (models.User, error) {
	var user models.User

	if username == "" {
		return user, nil
	}

	fbUsersRef := config.GetConfig().FbUsersRef
	fbUserRef := fbUsersRef.Child(username)

	err := fbUserRef.Get(context.Background(), &user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func CreateUser(user models.User) (models.User, error) {
	if user.Username == "" {
		return user, nil
	}

	fbUsersRef := config.GetConfig().FbUsersRef
	fbUserRef := fbUsersRef.Child(user.Username)
	fbUserRef.Set(context.Background(), user)

	return user, nil
}
