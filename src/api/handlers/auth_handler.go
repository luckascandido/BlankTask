package handlers

import (
	Requests "blanktask/src/api/requests"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

// função de captura de body
func (h *Handler) RegisterHandler(c echo.Context) error {
	//debug apagar depois
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	fmt.Println("Corpo da Requisição:", string(body))
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))
	payload := &Requests.RegisterUserRequest{}

	if err := c.Bind(payload); err != nil {
		return err
	}
	c.Logger().Info(payload)
	return c.String(http.StatusOK, "good request")
}
