package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alidevjimmy/go-echo-train/services"
	"github.com/alidevjimmy/go-echo-train/utils"
	"github.com/labstack/echo"
)

const (
	ErrorMessageInvalidInput = "input %s is not valid"
	ErrorInvalidInput        = "invalid_input"
)

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Wolcome to Dockerized App!!!")
}

func GetProductById(c echo.Context) error {
	id, err := convertInputStringToInt(c.Param("id"), "id")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	product, err := services.GetProduct(id)
	if err != nil {
		return c.JSON(err.Status, err)
	}

	return c.JSON(http.StatusOK, product)
}

func convertInputStringToInt(s string, key string) (int, *utils.ApiError) {
	value, err := strconv.Atoi(s)
	if err != nil {
		err := &utils.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf(ErrorMessageInvalidInput, key),
			Error:   ErrorInvalidInput,
		}
		return 0, err
	}
	return value, nil
}
