package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/IamNotUrKitty/gophermart/internal/domain/user"
	"github.com/labstack/echo/v4"
)

type RegisterUserDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *Handler) Register(c echo.Context) error {
	var userData RegisterUserDTO

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := json.Unmarshal(body, &userData); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	u := user.CreateUser(userData.Login, userData.Password)

	if err := h.repo.SaveUser(c.Request().Context(), *u); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, u)
}
