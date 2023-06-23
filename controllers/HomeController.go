package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HomePage(c echo.Context) error {
	data := map[string]interface{}{}
	return c.Render(http.StatusOK, "home.html", data)
}
