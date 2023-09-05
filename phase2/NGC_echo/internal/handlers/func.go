package handlers

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"ngc_echo/internal/model"
	"ngc_echo/pkg"
)

func (h *Handler) Register(c echo.Context) error {
	// bind reqbody
	var reqBody model.Users
	err := c.Bind(&reqBody)
	if err != nil {
		return errors.New("Failed bind")
	}

	if err = h.usecase.Register(c.Request().Context(), &reqBody); err != nil {
		return err
	}

	pkg.WriteResponseWithData(c, 201, "Success create", reqBody)
	return nil
}

func (h *Handler) Login(c echo.Context) error {
	// bind req body
	var reqBody model.Users
	err := c.Bind(&reqBody)
	if err != nil {
		return errors.New("Failed bind")
	}

	// find from db by username
	user, token, err := h.usecase.Login(c.Request().Context(), reqBody.Username, reqBody.Password)
	if err != nil {
		return err
	}

	// send to header response
	type ret struct {
		User      *model.Users `json:"user"`
		AuthToken string       `json:"auth_token"`
	}

	// response body
	pkg.WriteResponseWithData(c, 200, "Success login", ret{
		User:      user,
		AuthToken: token,
	})
	return nil
}

func (h *Handler) ProductsHandler(c echo.Context) error {
	user := c.Get("user").(model.Users)
	products, err := h.usecase.GetAllProducts(c.Request().Context())
	if err != nil {
		return err
	}

	pkg.WriteResponseWithData(c, 200, "logged in user "+user.Username, products)
	return nil
}

func (h *Handler) Buy(c echo.Context) error {
	// bind transaction
	var transaction model.Transactions
	err := c.Bind(&transaction)
	if err != nil {
		return errors.New("Failed bind")
	}

	user := c.Get("user").(model.Users)

	// find product
	trx, err := h.usecase.InsertTransaction(c.Request().Context(), &user, &transaction)
	if err != nil {
		return err
	}

	message := fmt.Sprintf("%s successfully purchased '%s'", user.Username)
	pkg.WriteResponseWithData(c, 201, message, trx)
	return nil
}
