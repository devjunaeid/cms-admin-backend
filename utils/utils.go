package utils

import (
	"github.com/devjunaeid/cms-admin-backend/types"
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
