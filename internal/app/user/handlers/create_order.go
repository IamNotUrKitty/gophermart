package handlers

import (
	"io"
	"net/http"

	"github.com/IamNotUrKitty/gophermart/internal/domain/order"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateOrder(c echo.Context) error {
	userID := c.Get("userId").(uuid.UUID)

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	o, err := order.CreateOrder(string(body))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := h.orderRepo.SaveOrder(c.Request().Context(), o, userID); err != nil {
		return c.String(http.StatusBadRequest, userID.String())
	}

	return c.String(http.StatusOK, string(body))
}
