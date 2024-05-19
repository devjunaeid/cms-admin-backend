package utils

import (
	"time"

	"github.com/devjunaeid/cms-admin-backend/config"
	"github.com/devjunaeid/cms-admin-backend/types"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJwt(userClaim types.DbLoginReqResponse) (string, error) {

	claims := jwt.MapClaims{
		"name":  userClaim.Name,
		"email": userClaim.Email,
		"role":  userClaim.Role,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	tokenString, err := token.SignedString([]byte(config.Env.JWTSecret))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
