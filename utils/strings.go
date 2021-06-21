package utils

import "golang.org/x/crypto/bcrypt"

func Hash(str string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
