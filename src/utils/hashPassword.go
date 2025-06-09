package utils

import (
	"golang.org/x/crypto/bcrypt"
)

const cost int = 10

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), cost)
}