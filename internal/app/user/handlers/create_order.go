package handlers

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateOrder(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, string(body))
}
