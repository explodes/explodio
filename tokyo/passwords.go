package tokyo

import (
	"github.com/explodes/explodio/stand"
	"golang.org/x/crypto/bcrypt"
)

var salt = stand.RequireEnv("PASSWORD_SALT")

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(salt+password), 4)
	return string(bytes), err
}

func CheckPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(salt+password))
	return err == nil
}
