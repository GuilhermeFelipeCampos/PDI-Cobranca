package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct {
	server *echo.Echo
}

func StartServer() *Server {
	e := echo.New()
	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
	return &Server{
		server: e,
	}
}
