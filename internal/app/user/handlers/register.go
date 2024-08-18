package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/IamNotUrKitty/gophermart/internal/domain/user"
	"github.com/IamNotUrKitty/gophermart/internal/echomiddleware"
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

	existedUser, err := h.userRepo.GetUser(c.Request().Context(), userData.Login)
	if existedUser != nil {
		return c.String(http.StatusConflict, errors.New("Already exist").Error())
	}

	u, err := user.CreateUser(userData.Login, userData.Password)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := h.userRepo.SaveUser(c.Request().Context(), *u); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	cookie, err := echomiddleware.GetUserToken(u)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, u)
}
