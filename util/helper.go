package util

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func JsonEncode(v interface{}) (string, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func GeneratePassword(userPassword string) (str string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	if err !=nil {
		return
	}
	str = string(bytes)
	return

}

func ValidatePassword(userPassword string, hashed string) (isOK bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		fmt.Println(err)
		return false, PasswordCheckError
	}
	return true, nil

}
