package handlers

import (
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/labstack/echo/v4"
)

type Error struct {
	Code    int
	Message string
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := err.Error()
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message.(string)
	}
	var errorResponse = Error{
		Code:    code,
		Message: message,
	}
	htmlFiles := []string{
		filepath.Join("./", "templates", "error.html"),
	}

	templ, _ := template.ParseFiles(htmlFiles...)

	templ.ExecuteTemplate(c.Response(), "error", errorResponse)
}
