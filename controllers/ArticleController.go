package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"myweb/models"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type Response struct {
	ErrorCode int         `json:"error_code" form:"error_code"`
	Message   string      `json:"message" form:"message"`
	Data      interface{} `json:"data"`
}

func ArticleIndex(c echo.Context) error {
	categories, err := models.GetAllCategories()
	if err != nil {
		return err
	}
	data := map[string]interface{}{
		"categories": categories,
	}
	return c.Render(http.StatusOK, "article_create.html", data)
}

func ArticlePage(c echo.Context) error {
	articles, err := models.GetAllArticles()
	if err != nil {
		return err
	}
	categories, err := models.GetAllCategories()
	if err != nil {
		return err
	}
	data := map[string]interface{}{
		"articles":   articles,
		"categories": categories,
	}
	return c.Render(http.StatusOK, "blog.html", data)
}

func ArticleCreate(c echo.Context) error {
	article := new(models.Article)
	c.Bind(article)
	response := new(Response)
	contentType := c.Request().Header.Get("Content-type")
	if contentType == "application/x-www-form-urlencoded" {
		if article.CreateArticle() != nil {
			response.ErrorCode = 10
			response.Message = "Gagal create data user"
		} else {
			response.ErrorCode = 0
			response.Message = "Sukses create data user"
			response.Data = *article
		}
	}
	return c.Redirect(http.StatusMovedPermanently, "/")
}

func ArticleUploadHandler(c echo.Context) error {
	// Read File

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Make folder
	err = os.MkdirAll("./images", os.ModePerm)
	if err != nil {
		return err
	}

	// Create new file
	filename := fmt.Sprintf("/images/%d%s", time.Now().UnixNano(), filepath.Ext(file.Filename))
	dst, err := os.Create("." + filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy the uploaded
	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	response := map[string]string{"location": filename}
	// response, _ := json.Marshal(map[string]string{"location": filename})
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(201)
	return c.JSON(http.StatusCreated, response)

}

func ArticleServeImages(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fs := http.StripPrefix("/images/", http.FileServer(http.Dir("./images")))
	fs.ServeHTTP(w, r)
}

func ArticleEdit(c echo.Context) error {
	var article models.Article
	var err error
	id, _ := strconv.Atoi(c.Param("id"))
	article, err = models.GetOneArticle(id)
	if err != nil {
		return nil
	}
	data := map[string]interface{}{
		"articleTitle":   article.Title,
		"articleContent": article.Content,
		"articleID":      article.ID,
	}
	return c.Render(http.StatusOK, "edit_article.html", data)
}

func ArticleUpdate(c echo.Context) error {
	article := new(models.Article)
	c.Bind(article)
	id, _ := strconv.Atoi(c.Param("id"))
	response := new(Response)
	if article.UpdateArticle(id) != nil {
		response.ErrorCode = 10
		response.Message = "Gagal update article"
	} else {
		response.ErrorCode = 200
		response.Message = "Update article berhasil"
		response.Data = *article
	}
	return c.Redirect(http.StatusMovedPermanently, "/articles")
}

func ArticleDelete(c echo.Context) error {
	var article models.Article
	var err error
	id, _ := strconv.Atoi(c.Param("id"))
	article, err = models.GetOneArticle(id)
	if err != nil {
		return nil
	}
	response := new(Response)
	if article.DeleteArticle() != nil {
		response.ErrorCode = 10
		response.Message = "Gagal delete article"
	} else {
		response.ErrorCode = 200
		response.Message = "Delete article berhasil"
	}
	return c.Redirect(http.StatusMovedPermanently, "/articles")
}

func ArticleSearch(c echo.Context) error {
	articles, err := models.GetAll(c.QueryParam("keywords"))
	if err != nil {
		return err
	}
	categories, err := models.GetAllCategories()
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"articles":   articles,
		"categories": categories,
	}

	return c.Render(http.StatusOK, "blog.html", data)
}

func ArticleSearchByCategory(c echo.Context) error {
	category, err := models.GetOneCategory(c.QueryParam("keywords"))
	if err != nil {
		return err
	}
	categories, err := models.GetAllCategories()
	if err != nil {
		return err
	}
	data := map[string]interface{}{
		"category":   category,
		"categories": categories,
	}

	return c.Render(http.StatusOK, "blog.html", data)
}
