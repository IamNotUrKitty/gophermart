package user

import (
	"github.com/IamNotUrKitty/gophermart/internal/app/user/handlers"
	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo, userRepo handlers.UserRepository, orderRepo handlers.OrderRepository) {
	handler := handlers.NewHandler("test", userRepo, orderRepo)

	e.POST("/api/user/register", handler.Register)
	e.POST("/api/user/login", handler.Login)
	e.POST("/api/user/orders", handler.CreateOrder)
	e.GET("/api/user/orders", handler.Orders)
	e.GET("/api/user/balance", handler.Balance)
	e.POST("/api/user/withdraw", handler.Withdraw)
	e.GET("/api/user/withdrawals", handler.Withdrawals)
}
