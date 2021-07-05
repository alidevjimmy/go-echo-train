package utils

import (
	"fmt"
	"net/http"
)

type ApiError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func ErrorNotFound(itemName string) *ApiError {
	return &ApiError{
		Status:  http.StatusNotFound,
		Message: fmt.Sprintf("%s not found", itemName),
		Error:   "not_found",
	}
}
