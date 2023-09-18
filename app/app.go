package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
	controller "myweb/controllers"
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

func StartApp() {

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
	//e.Static("/", "images")

	//Routes
	e.GET("/user/signin", controller.SignInForm()).Name = "userSignInForm"
	e.POST("/user/signin", controller.SignIn())
	PageHandler(e)
	ArticleHandler(e)
	AdminHandler(e)

	//End Routes

	e.Logger.Fatal(e.Start(":1323"))
}
