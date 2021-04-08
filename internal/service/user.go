package service

import (
	"errors"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/umarkotak/lul-er_BE/internal/models"
	"github.com/umarkotak/lul-er_BE/internal/repository"
	"github.com/umarkotak/lul-er_BE/internal/utils"
)

func Register(reqUser models.User) (models.User, error) {
	user, err := repository.GetUserByUsername(reqUser.Username)

	if err != nil {
		log.Printf("Error GetUserByUsername %v", err)
		return user, err
	}

	if user.Username != "" {
		return user, errors.New("username already taken")
	}

	userTokenClaim := jwt.StandardClaims{
		Issuer: reqUser.Username,
	}
	reqUser.AuthToken, err = utils.EncodeToken(userTokenClaim)
	if err != nil {
		log.Printf("Error EncodeToken %v", err)
		return user, err
	}

	pw, err := utils.EncodePassword(reqUser.Password)

	if err != nil {
		log.Printf("Error EncodeToken %v", err)
		return user, err
	}

	newUser := models.User{

		Username:         reqUser.Username,
		HashPassword:     pw,
		Email:            reqUser.Email,
		AuthToken:        reqUser.AuthToken,
		ActiveGameRoomID: "",
	}

	_, err = repository.CreateUser(newUser)
	if err != nil {
		log.Printf("Error CreateUser %v", err)
		return user, err
	}

	return newUser, nil
}

func Login(reqUser models.User) (models.User, error) {

	user, err := repository.GetUserByUsername(reqUser.Username)

	if err != nil {
		log.Printf("Error GetUserByUsername %v", err)
		return user, err
	}

	if user.Username != reqUser.Username {
		return user, errors.New("username is wrong")
	}

	isPassword := utils.DecodePassword(reqUser.Password, user.HashPassword)

	if !isPassword {
		return user, errors.New("Invalid Password")
	}

	userData := models.User{

		Username:  user.Username,
		AuthToken: user.AuthToken,
	}

	return userData, nil

}
