package internal

import (
	"net/http"

	"github.com/IamNotUrKitty/gophermart/internal/app/user"
	"github.com/IamNotUrKitty/gophermart/internal/db"
	"github.com/IamNotUrKitty/gophermart/internal/echomiddleware"
	"github.com/IamNotUrKitty/gophermart/internal/infrastructure/orders"
	userInfra "github.com/IamNotUrKitty/gophermart/internal/infrastructure/user"
	"github.com/labstack/echo/v4"
)

func NewServer() (*echo.Echo, error) {
	e := echo.New()

	e.Use(echomiddleware.InitJWTMiddleware())

	pool, err := db.NewConnectionPool("postgres://postgres:123@localhost:5432/gophermart?sslmode=disable")
	if err != nil {
		return nil, err
	}

	userRepo, err := userInfra.NewPostgressRepo(pool)
	if err != nil {
		return nil, err
	}

	orderRepo, err := orders.NewPostgressRepo(pool)
	if err != nil {
		return nil, err
	}

	user.Setup(e, userRepo, orderRepo)

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	return e, nil
}
