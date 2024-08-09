package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateOrder(c echo.Context) error {
	return c.String(http.StatusOK, "create order")
}
