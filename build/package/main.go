package main

import (
	"PDI-COBRANCA/build/package/handlers"
	"PDI-COBRANCA/build/package/server"

	_ "github.com/lib/pq"
)

func main() {

	//find all
	server.E.GET("/products", handlers.GetAll)
	//find by email
	server.E.GET("/product", handlers.GetByEmail)
	// record products
	server.E.POST("/products", handlers.Create)
	// update product
	server.E.PUT("/products/:id", handlers.Update)
	//delete by id
	server.E.DELETE("/products/:id", handlers.Delete)

	server.StartServer()

}
