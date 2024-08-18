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

type LoginUserDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *Handler) Login(c echo.Context) error {
	var userData RegisterUserDTO

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := json.Unmarshal(body, &userData); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	u, err := user.CreateUser(userData.Login, userData.Password)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	existedUser, err := h.userRepo.GetUser(c.Request().Context(), u.Username())
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	if existedUser == nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	if existedUser.PasswordHash() == u.PasswordHash() {
		cookie, err := echomiddleware.GetUserToken(existedUser)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		c.SetCookie(cookie)
	} else {
		return c.String(http.StatusUnauthorized, errors.New("Неверная пара логин/пароль").Error())
	}

	return c.JSON(http.StatusOK, u)
}
