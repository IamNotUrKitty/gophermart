package internal

import (
	"net/http"

	"github.com/IamNotUrKitty/gophermart/internal/app/user"
	userInfra "github.com/IamNotUrKitty/gophermart/internal/infrastructure/user"
	"github.com/labstack/echo/v4"
)

func NewServer() (*echo.Echo, error) {

	e := echo.New()

	userRepo, err := userInfra.NewPostgressRepo("postgres://postgres:123@localhost:5432/gophermart?sslmode=disable")
	if err != nil {
		return nil, err
	}

	user.Setup(e, userRepo)

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	return e, nil
}
