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
		return c.String(http.StatusBadRequest, "неверный формат запроса")
	}

	o, err := order.CreateOrder(string(body))
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, "неверный формат номера заказа")
	}

	eo, err := h.orderRepo.GetOrder(c.Request().Context(), o)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if eo != nil {
		if eo.UserID == userID {
			return c.String(http.StatusOK, "номер заказа уже был загружен этим пользователем")
		}
		return c.String(http.StatusConflict, "номер заказа уже был загружен другим пользователем")
	}

	if err := h.orderRepo.SaveOrder(c.Request().Context(), o, userID); err != nil {
		return c.String(http.StatusInternalServerError, userID.String())
	}

	return c.String(http.StatusAccepted, "новый номер заказа принят в обработку")
}
