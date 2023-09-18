package app

import (
	"github.com/labstack/echo/v4"
	controller "myweb/controllers"
)

func PageHandler(e *echo.Echo) {
	e.GET("/", controller.HomePage)
	e.GET("/articles", controller.ArticlePage)
	e.GET("/articles_create", controller.ArticleIndex)

}
