package handlers

import (
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/Danila331/HACH-T1/app/models"
	"github.com/labstack/echo/v4"
)

func MakerPage(c echo.Context) error {
	var node = models.Node{}
	nodes, err := node.ReadAll()
	fmt.Println(nodes)
	if err != nil {
		return err
	}

	htmlFiles := []string{
		filepath.Join("./", "templates", "maker.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "maker", nodes)
	return nil
}

func MakerAddPage(c echo.Context) error {
	htmlFiles := []string{
		filepath.Join("./", "templates", "nodes.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "nodes", nil)
	return nil
}

func MakerAddSubmit(c echo.Context) error {
	return nil
}
