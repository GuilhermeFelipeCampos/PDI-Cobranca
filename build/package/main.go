package main

import (
	"PDI-COBRANCA/build/package/handlers"
	"PDI-COBRANCA/build/package/server"

	_ "github.com/lib/pq"
)

func main() {
	//pudim := handlers.NewHandler(5)
	//find all
	server.E.GET("/users", handlers.GetAll)
	//find by email
	server.E.GET("/user", handlers.GetByEmail)
	// record products
	server.E.POST("/user", handlers.Create)
	// update product
	server.E.PUT("/user/:id", handlers.Update)
	//delete by id
	server.E.DELETE("/user/:id", handlers.Delete)

	server.StartServer()

}
