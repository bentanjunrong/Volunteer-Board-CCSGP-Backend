package utils

import "golang.org/x/crypto/bcrypt"

func Hash(str string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(str), 14)
	return string(bytes)
}
