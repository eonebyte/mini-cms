package app

import (
	"github.com/labstack/echo/v4"
	controller "myweb/controllers"
	"net/http"
)

func ArticleHandler(e *echo.Echo) {
	e.GET("/articles/search", controller.ArticleSearch)
	e.GET("/articles/search_by_category", controller.ArticleSearchByCategory)
	e.POST("/articles_post", controller.ArticleCreate)
	e.GET("/articles/:id", controller.ArticleEdit)
	e.PUT("/article_update/:id", controller.ArticleUpdate)
	e.DELETE("/article_delete/:id", controller.ArticleDelete)
	e.POST("/upload", controller.ArticleUploadHandler)
	e.GET("/images/*", echo.WrapHandler(http.HandlerFunc(controller.ArticleServeImages)))
}
