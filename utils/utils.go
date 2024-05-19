package utils

import (
	"github.com/devjunaeid/cms-admin-backend/types"
	"golang.org/x/crypto/bcrypt"
)

// Create Error Response JSON.
func CreateErrorRes(msg string, status int) types.ErrorDefault {
	return types.ErrorDefault{
		Status: status,
		Err:    msg,
	}
}

// Create Operation Success Response JSON.
func CreateSuccessRes(msg string, status int) types.SuccessDefault {
	return types.SuccessDefault{
		Status: status,
		Msg:    msg,
	}
}

func JwtTokenRes(token string, status int) types.JwtTokenRes {
	return types.JwtTokenRes{
		Token:  token,
		Status: status,
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
