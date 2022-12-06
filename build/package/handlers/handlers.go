package handlers

import (
	"PDI-COBRANCA/build/package/model"
	"PDI-COBRANCA/build/package/repository"
	"PDI-COBRANCA/build/package/server"
	"fmt"

	"net/http"
	"strconv"

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

type ProductValidator struct {
	validator *validator.Validate
}

func (p *ProductValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

var (
	products = []map[int]string{{1: "moblies"}, {2: "tvs"}, {3: "laptops"}}

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
	server.E.Validator = &ProductValidator{validator: v}
	if err := c.Bind(&uEmail); err != nil {
		return err
	}
	if err := c.Validate(uEmail); err != nil {
		return c.JSON(http.StatusBadRequest, "Error in request")
	}
	email := fmt.Sprintf("%v", uEmail)
	fmt.Println("handler: " + email)
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

	server.E.Validator = &ProductValidator{validator: v}
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
	var product map[int]string
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	for _, p := range products {
		for k := range p {
			if pID == k {
				product = p
			}
		}
	}
	if product == nil {
		return c.JSON(http.StatusNotFound, "product not found")
	}
	type body struct {
		Name string `json:"product_name" validate:"required,min=4"`
	}
	var reqBody body
	server.E.Validator = &ProductValidator{validator: v}
	if err := c.Bind(&reqBody); err != nil {
		return err
	}
	if err := c.Validate(reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, "Error in request")
	}
	product[pID] = reqBody.Name
	return c.JSON(http.StatusOK, product)
}

func Delete(c echo.Context) error {
	var product map[int]string
	var index int
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	for i, p := range products {
		for k := range p {

			if pID == k {
				product = p
				index = i
			}
		}
	}
	if product == nil {
		return c.JSON(http.StatusNotFound, "product not found")
	}
	splice := func(s []map[int]string, index int) []map[int]string {
		return append(s[:index], s[index+1:]...)
	}
	products = splice(products, index)
	return c.JSON(http.StatusOK, product)
}
