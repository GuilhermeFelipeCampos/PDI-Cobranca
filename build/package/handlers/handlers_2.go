package handlers

import (
	"PDI-COBRANCA/build/package/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
	repository repository.RepositoryInterface
}

func NewHandler(repository repository.RepositoryInterface) handler {

	return handler{
		repository: repository,
	}
}

func (h *handler) GetAll(c echo.Context) error {

	resp, err := h.repository.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error")
	}
	return c.JSON(http.StatusOK, resp)
}
