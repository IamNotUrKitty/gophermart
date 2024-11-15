package handlers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Orders(c echo.Context) error {
	userID := c.Get("userId").(uuid.UUID)

	fmt.Println(userID)

	orders, err := h.orderRepo.GetOrdersByUserID(c.Request().Context(), userID)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, orders)
}
