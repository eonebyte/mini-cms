package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"myweb/auth"
	"myweb/models"
	"net/http"
)

func Admin() echo.HandlerFunc {
	return func(c echo.Context) error {
		//Get user cookie
		userCookie, err := c.Cookie("user")
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, fmt.Sprintf("Hi, %s you have access!", userCookie.Value))
	}
}

func SignInForm() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "signIn.html", map[string]interface{}{})
	}

}

func SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		storedUser := models.LoadTestUser()
		req := new(models.User)

		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(req.Password)); err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Password is inccored")
		}

		err := auth.GenerateTokenAndSetCookie(storedUser, c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("error : %s", err.Error()))
		}
		return c.Redirect(http.StatusMovedPermanently, "/admin")
	}
}
