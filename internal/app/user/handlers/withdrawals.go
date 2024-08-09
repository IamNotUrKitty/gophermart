package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Withdrawals(c echo.Context) error {
	return c.String(http.StatusOK, "Withdrawals")
}
