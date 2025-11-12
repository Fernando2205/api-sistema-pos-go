package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword genera un hash bcrypt de la contraseña
func HashPassword(password string) (string, error) {
	// bcrypt.DefaultCost es 10, suficiente para la mayoría de casos
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}
