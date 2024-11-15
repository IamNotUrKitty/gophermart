package echomiddleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/IamNotUrKitty/gophermart/internal/domain/user"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type CustomClaims struct {
	UserID uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}

const Secret = "secret"
const CookieName = "user"

func InitJWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		Skipper: func(c echo.Context) bool {
			path := c.Request().URL.Path

			return strings.Contains(path, "login") || strings.Contains(path, "register")
		},
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(CustomClaims)
		},
		SuccessHandler: func(c echo.Context) {
			user := c.Get("user").(*jwt.Token)

			claims := user.Claims.(*CustomClaims)

			c.Set("userId", claims.UserID)
		},
		TokenLookup: "cookie:" + CookieName,
		SigningKey:  []byte(Secret),
	})
}

func GetUserToken(u *user.User) (*http.Cookie, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 3)),
		},
		UserID: u.ID(),
	})

	tokenString, err := token.SignedString([]byte(Secret))
	if err != nil {
		return nil, err
	}

	cookie := http.Cookie{
		Name:    CookieName,
		Value:   tokenString,
		Expires: time.Now().Add(3 * time.Hour),
	}

	return &cookie, nil
}
