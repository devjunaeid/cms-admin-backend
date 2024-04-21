package utils

import (
	"github.com/devjunaeid/cms-admin-backend/types"
)

func CreateErrorRes(msg string, status int) types.ErrorDefault {
	return types.ErrorDefault{
		Status: status,
		Err:    msg,
	}
}
