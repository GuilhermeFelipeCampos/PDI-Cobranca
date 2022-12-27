package handlers

import (
	"PDI-COBRANCA/build/package/model"
	"PDI-COBRANCA/build/package/repository"
	"PDI-COBRANCA/build/package/server"
	"fmt"

	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type HandlerInterface interface {
	GetAll(c echo.Context) error
	GetById(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type UserValidator struct {
	validator *validator.Validate
}

func (p *UserValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

var (
	v = validator.New()
)

func GetAll(c echo.Context) error {
	resp, err := repository.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error")
	}
	return c.JSON(http.StatusOK, resp)
}

func GetByEmail(c echo.Context) error {
	type emailRequest struct {
		Email string `query:"email" validate:"required,email"`
	}
	var uEmail emailRequest
	server.E.Validator = &UserValidator{validator: v}
	if err := c.Bind(&uEmail); err != nil {
		return err
	}
	if err := c.Validate(uEmail); err != nil {
		return c.JSON(http.StatusBadRequest, "Error in request")
	}
	email := fmt.Sprintf("%v", uEmail)
	resp, err := repository.GetUserByEmail(email)
	if err != nil {
		return c.JSON(http.StatusExpectationFailed, "Erro na requisição com o banco de dados")
	}
	return c.JSON(http.StatusOK, resp)
}

func Create(c echo.Context) error {
	type body struct {
		Name    string `json:"name" validate:"required"`
		Email   string `json:"email" validate:"required"`
		Keyword string `json:"keyword" validate:"required"`
	}
	var reqBody body

	server.E.Validator = &UserValidator{validator: v}
	if err := c.Bind(&reqBody); err != nil {
		return err
	}
	if err := c.Validate(reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, "Error in request")
	}
	users := model.Users{
		Name:    reqBody.Name,
		Email:   reqBody.Email,
		Keyword: reqBody.Keyword,
	}

	resp, err := repository.InsertUsers(users)
	if err != nil {
		return c.JSON(http.StatusExpectationFailed, "Erro na requisição com o banco de dados")
	}
	return c.JSON(http.StatusCreated, resp)
}

func Update(c echo.Context) error {
	userID := c.Param("id")

	type body struct {
		Name    string `json:"name" validate:"required"`
		Email   string `json:"email" validate:"required"`
		Keyword string `json:"keyword" validate:"required"`
	}
	var reqBody body

	server.E.Validator = &UserValidator{validator: v}
	if err := c.Bind(&reqBody); err != nil {
		return err
	}
	if err := c.Validate(reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, "Error in request")
	}
	users := model.Users{
		Id:      userID,
		Name:    reqBody.Name,
		Email:   reqBody.Email,
		Keyword: reqBody.Keyword,
	}
	resp, err := repository.UpdateUser(users)
	if err != nil {
		return c.JSON(http.StatusExpectationFailed, "Erro na requisição com o banco de dados")
	}
	return c.JSON(http.StatusOK, resp)
}

func Delete(c echo.Context) error {
	userID := c.Param("id")
	resp, err := repository.DeleteUser(userID)
	if err != nil {
		return c.JSON(http.StatusExpectationFailed, "Erro na requisição com o banco de dados")
	}
	return c.JSON(http.StatusOK, resp)
}
