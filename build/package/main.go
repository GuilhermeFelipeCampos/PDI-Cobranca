package main

import (
	"PDI-COBRANCA/build/package/handlers"
	"PDI-COBRANCA/build/package/server"

	"github.com/labstack/echo/v4"
)

func main() {
	server.StartServer()
	e := echo.New()
	//find all
	e.GET("/products", handlers.GetAll)
	//find by id
	e.GET("/products/:id", handlers.GetById)
	// record products
	e.POST("/products", handlers.Create)
	// update product
	e.PUT("/products/:id", handlers.Update)
	//delete by id
	e.DELETE("/products/:id", handlers.Delete)

}
