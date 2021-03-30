package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

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

	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: reqUser.Username,
	})

	token, err := claim.SignedString([]byte("secret"))

	if err != nil {
		log.Errorf(context.Background(), "Error Token %v", err)
		return "", err
	}
	return token, nil

}

func DecodeToken(claimToken string) (string, error) {

	clearToken := strings.Split(claimToken, " ")[1]

	token, err := jwt.ParseWithClaims(clearToken, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	claims := token.Claims.(*jwt.StandardClaims)
	claimUser := claims.Issuer

	return claimUser, nil
}
