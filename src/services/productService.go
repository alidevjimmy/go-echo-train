package services

import (
	"github.com/alidevjimmy/go-echo-train/database"
	"github.com/alidevjimmy/go-echo-train/domains"
	"github.com/alidevjimmy/go-echo-train/utils"
)

func GetProduct(id int) (domains.Product, *utils.ApiError) {
	for _, product := range database.Products {
		if product.Id == id {
			return product, nil
		}
	}
	return domains.Product{}, utils.ErrorNotFound("product")
}
