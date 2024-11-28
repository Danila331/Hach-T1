package handlers

import (
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/Danila331/HACH-T1/app/models"
	"github.com/Danila331/HACH-T1/app/pkg"
	"github.com/labstack/echo/v4"
)

func DeleteFileSubmit(c echo.Context) error {
	id, _ := strconv.Atoi(c.FormValue("id"))
	var file = models.File{
		ID: id,
	}
	file, err := file.ReadByID()
	if err != nil {
		return err
	}
	err = pkg.S3DeleteFile(file.Name)
	if err != nil {
		return err
	}
	err = file.Delete()
	if err != nil {
		return err
	}
	htmlFiles := []string{filepath.Join("./", "templates", "submit", "delete", "file_submit.html")}
	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}
	templ.ExecuteTemplate(c.Response(), "file_submit", nil)
	return nil
}

func DeleteDatabaseSubmit(c echo.Context) error {
	id, _ := strconv.Atoi(c.FormValue("id"))
	var database = models.Database{
		ID: id,
	}
	err := database.Delete()
	if err != nil {
		return err
	}
	htmlFiles := []string{filepath.Join("./", "templates", "submit", "delete", "database_submit.html")}
	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}
	templ.ExecuteTemplate(c.Response(), "database_submit", nil)
	return nil
}

func DeleteWebsiteSubmit(c echo.Context) error {
	id, _ := strconv.Atoi(c.FormValue("id"))
	var website = models.Website{
		ID: id,
	}
	err := website.Delete()
	if err != nil {
		return err
	}
	htmlFiles := []string{filepath.Join("./", "templates", "submit", "delete", "website_submit.html")}
	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}
	templ.ExecuteTemplate(c.Response(), "website_submit", nil)
	return nil
}

func MakerDeleteSubmit(c echo.Context) error {
	id, _ := strconv.Atoi(c.FormValue("id"))
	var node = models.Node{
		ID: id,
	}
	err := node.Delete()
	if err != nil {
		return err
	}
	htmlFiles := []string{filepath.Join("./", "templates", "submit", "delete", "node_submit.html")}
	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}
	templ.ExecuteTemplate(c.Response(), "node_submit", nil)
	return nil
}
