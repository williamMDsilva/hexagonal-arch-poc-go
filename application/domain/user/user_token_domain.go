package domain

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/williamMDsilva/hexagonal-arch-poc-go.git/application/domain/user/constants"
	http_errors "github.com/williamMDsilva/hexagonal-arch-poc-go.git/configuration/errors"
)

func (ud *UserDomain) GenerateToken() (string, *http_errors.HttpErr) {
	secret := os.Getenv(constants.JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":    ud.Id,
		"email": ud.Email,
		"name":  ud.Name,
		"age":   ud.Age,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // token expiration
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", http_errors.NewInternalServerError(
			fmt.Sprintf("error trying to generate jwt token, err=%s", err.Error()))
	}

	return tokenString, nil
}
