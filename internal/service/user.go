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

func Register(reqUser models.User) (models.User, error) {
	user, err := repository.GetUserByUsername(reqUser.Username)

	if err != nil {
		log.Errorf(context.Background(), "Error GetUserByUsername %v", err)
		return user, err
	}

	if user.Username != "" {
		return user, errors.New("username already taken")
	}

	reqUser.AuthToken, _ = encodeToken(reqUser)

	newUser := models.User{
		Username:         reqUser.Username,
		Password:         reqUser.Password,
		Email:            reqUser.Email,
		AuthToken:        reqUser.AuthToken,
		ActiveGameRoomID: "",
	}

	_, err = repository.CreateUser(newUser)
	if err != nil {
		log.Errorf(context.Background(), "Error CreateUser %v", err)
		return user, err
	}

	return newUser, nil
}

func Login(reqUser models.User) (models.User, error) {

	user, err := repository.GetUserByUsername(reqUser.Username)

	if err != nil {
		log.Errorf(context.Background(), "Error GetUserByUsername %v", err)
		return user, err
	}

	if user.Username != reqUser.Username || user.Password != reqUser.Password {

		return user, errors.New("username or password is wrong")
	}

	userData := models.User{

		Username:  user.Username,
		AuthToken: user.AuthToken,
	}

	return userData, nil

}

func encodeToken(reqUser models.User) (string, error) {

	lulErJwtSecret := os.Getenv("GO_LULER_JWT_SECRET")
	tokenData := jwt.MapClaims{}
	tokenData["username"] = reqUser.Username
	claimToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenData)
	token, err := claimToken.SignedString([]byte(lulErJwtSecret))
	if err != nil {

		log.Errorf(context.Background(), "Error Token %v", err)
		return "", err
	}
	return token, nil

}
