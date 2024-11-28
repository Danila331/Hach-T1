package handlers

import (
	"path/filepath"
	"text/template"

	"github.com/Danila331/HACH-T1/app/models"
	"github.com/labstack/echo/v4"
)

func MakerTrigerPage(c echo.Context) error {
	htmlFiles := []string{
		filepath.Join("./", "templates", "node", "triger.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "triger", nil)
	return nil
}

func MakerTrigerPageSubmit(c echo.Context) error {
	text := c.FormValue("text")
	ShouldToDo := c.FormValue("shoudltodo")
	var node = models.Node{
		Type:       "triger",
		Text:       text,
		ShouldToDo: ShouldToDo,
	}
	err := node.Create()
	if err != nil {
		return err
	}
	htmlFiles := []string{
		filepath.Join("./", "templates", "submit", "node_submit.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "node_submit", nil)
	return nil
}

func MakerStartPage(c echo.Context) error {
	htmlFiles := []string{
		filepath.Join("./", "templates", "node", "start.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "start", nil)
	return nil
}

func MakerStartPageSubmit(c echo.Context) error {
	text := c.FormValue("text")
	var node = models.Node{
		Type:       "start",
		Text:       text,
		ShouldToDo: "",
	}
	err := node.Create()
	if err != nil {
		return err
	}
	htmlFiles := []string{
		filepath.Join("./", "templates", "submit", "node_submit.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "node_submit", nil)
	return nil
}

func MakerEndPage(c echo.Context) error {
	htmlFiles := []string{
		filepath.Join("./", "templates", "node", "end.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "end", nil)
	return nil
}

func MakerEndPageSubmit(c echo.Context) error {
	text := c.FormValue("text")
	var node = models.Node{
		Type:       "end",
		Text:       text,
		ShouldToDo: "",
	}
	err := node.Create()
	if err != nil {
		return err
	}
	htmlFiles := []string{
		filepath.Join("./", "templates", "submit", "node_submit.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "node_submit", nil)
	return nil
}

func MakerErrorPage(c echo.Context) error {
	htmlFiles := []string{
		filepath.Join("./", "templates", "node", "error.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "error", nil)
	return nil
}

func MakerErrorPageSubmit(c echo.Context) error {
	text := c.FormValue("text")
	var node = models.Node{
		Type:       "error",
		Text:       text,
		ShouldToDo: "",
	}
	err := node.Create()
	if err != nil {
		return err
	}
	htmlFiles := []string{
		filepath.Join("./", "templates", "submit", "node_submit.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "node_submit", nil)
	return nil
}
