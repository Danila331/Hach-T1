package handlers

import (
	"path/filepath"
	"text/template"

	"github.com/Danila331/HACH-T1/app/models"
	"github.com/labstack/echo/v4"
)

func AddWebsitePage(c echo.Context) error {
	htmlFiles := []string{
		filepath.Join("./", "templates", "website.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "website", nil)
	return nil
}

func AddWebsiteSubmit(c echo.Context) error {
	name := c.FormValue("name")
	url := c.FormValue("url")
	var website = models.Website{
		Name: name,
		URL:  url,
	}
	err := website.Create()
	if err != nil {
		return err
	}
	htmlFiles := []string{
		filepath.Join("./", "templates", "submit", "website_submit.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "website_submit", nil)
	return nil
}
