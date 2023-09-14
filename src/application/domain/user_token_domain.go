package domain

import (
	"fmt"
	"os"
	"time"

	"github.com/Doehnert/crud-hexa/src/application/constants"
	"github.com/Doehnert/crud-hexa/src/configuration/rest_errors"
	"github.com/golang-jwt/jwt/v4"
)

func (ud *UserDomain) GenerateToken() (string, *rest_errors.RestErr) {
	secret := os.Getenv(constants.JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":    ud.Id,
		"email": ud.Email,
		"name":  ud.Name,
		"age":   ud.Age,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", rest_errors.NewInternalServerError(
			fmt.Sprintf("error trying to generate jwt token, err=%s", err.Error()),
		)
	}

	return tokenString, nil
}
