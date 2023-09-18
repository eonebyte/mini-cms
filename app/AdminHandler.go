package app

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"myweb/auth"
	controller "myweb/controllers"
	"net/http"
)

func AdminHandler(e *echo.Echo) {
	adminGroup := e.Group("/admin")
	adminGroup.Use(echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.Claims)
		},
		SigningKey:  []byte(auth.GetJwtSecret()),
		TokenLookup: "cookie:access-token",
		ErrorHandler: func(c echo.Context, err error) error {
			return echo.NewHTTPError(http.StatusUnauthorized, "Silahkan Login Dulu")
		},
	}))
	//adminGroup.Use(echojwt.JWT([]byte(auth.GetJwtSecret())))

	adminGroup.Use(auth.TokenRefreshMiddleware)

	adminGroup.GET("", controller.Admin())
}
