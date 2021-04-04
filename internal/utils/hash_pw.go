package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func EncodePassword(userPassword string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	if err != nil {

		fmt.Println(err)
	}

	fmt.Println(hash)
	return hash, err

}

func DecodePassword(userPassword string, hashpassword string) (bool, error) {

	if err := bcrypt.CompareHashAndPassword(hashpassword, []byte(userPassword)); err != nil {

		return false, err
	}

	return true, err
}
