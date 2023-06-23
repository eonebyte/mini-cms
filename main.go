package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
	config "myweb/config"
	controller "myweb/controllers"
	"net/http"
)

type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func catch(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

type Response struct {
	Code    int         `json:"code" form:"code"`
	Message string      `json:"message" form:"message"`
	Data    interface{} `json:"data"`
}

func main() {

	config.ConnectDB()
	e := echo.New()
	e.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
		Getter: middleware.MethodFromForm("_method"),
	}))
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = renderer
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "assets")
	e.GET("/", controller.HomePage)
	e.GET("/articles/search", controller.ArticleSearch)
	e.GET("/articles/search_by_category", controller.ArticleSearchByCategory)
	// CRUD
	e.POST("/articles_post", controller.ArticleCreate)
	e.GET("/articles_create", controller.ArticleIndex)
	e.GET("/articles", controller.ArticlePage)
	e.GET("/articles/:id", controller.ArticleEdit)
	e.PUT("/article_update/:id", controller.ArticleUpdate)
	e.DELETE("/article_delete/:id", controller.ArticleDelete)
	// END CRUD
	e.POST("/upload", controller.ArticleUploadHandler)
	e.GET("/images/*", echo.WrapHandler(http.HandlerFunc(controller.ArticleServeImages)))

	e.Logger.Fatal(e.Start(":1323"))
}
