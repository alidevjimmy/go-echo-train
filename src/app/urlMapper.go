package app

import (
	"github.com/alidevjimmy/go-echo-train/controllers"
	"github.com/alidevjimmy/go-echo-train/middlewares"
)

func router() {
	e.GET("/", controllers.Index)
	e.GET("/product/:id", controllers.GetProductById, middlewares.OnlyUsers)
}