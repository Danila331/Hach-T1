package handlers

import (
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/Danila331/HACH-T1/app/models"
	"github.com/labstack/echo/v4"
)

func EditDataBasePage(c echo.Context) error {
	dbId := c.QueryParam("id")
	DatabaseID, _ := strconv.Atoi(dbId)
	var database models.Database = models.Database{
		ID: DatabaseID,
	}
	database, err := database.ReadByID()
	if err != nil {
		return err
	}
	htmlFiles := []string{filepath.Join("./", "templates", "edit", "database_edit.html")}
	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}
	templ.ExecuteTemplate(c.Response(), "database_edit", database)
	return nil
}

func EditDataBasePageSubmit(c echo.Context) error {
	id, _ := strconv.Atoi(c.FormValue("id"))
	host := c.FormValue("host")
	port := c.FormValue("port")
	username := c.FormValue("user")
	password := c.FormValue("password")
	databaseName := c.FormValue("dbname")
	var database = models.Database{
		ID:           id,
		Host:         host,
		Port:         port,
		UserName:     username,
		Password:     password,
		DatabaseName: databaseName,
	}
	err := database.Update()
	if err != nil {
		return err
	}
	htmlFiles := []string{filepath.Join("./", "templates", "submit", "edit", "database_submit.html")}
	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}
	templ.ExecuteTemplate(c.Response(), "database_submit", database)
	return nil
}

func EditWebsitePage(c echo.Context) error {
	WebsiteID, _ := strconv.Atoi(c.QueryParam("id"))
	var website models.Website = models.Website{
		ID: WebsiteID,
	}
	website, err := website.ReadByID()
	if err != nil {
		return err
	}
	htmlFiles := []string{
		filepath.Join("./", "templates", "edit", "website_edit.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "website_edit", website)
	return nil
}

func EditWebsitePageSubmit(c echo.Context) error {
	id, _ := strconv.Atoi(c.FormValue("id"))
	name := c.FormValue("name")
	url := c.FormValue("url")
	var website = models.Website{
		ID:   id,
		Name: name,
		URL:  url,
	}
	err := website.Update()
	if err != nil {
		return err
	}
	htmlFiles := []string{
		filepath.Join("./", "templates", "submit", "edit", "website_submit.html"),
	}
	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}
	templ.ExecuteTemplate(c.Response(), "website_submit", nil)
	return nil
}

func MakerEditPage(c echo.Context) error {
	NodeId, _ := strconv.Atoi(c.QueryParam("id"))
	var node models.Node = models.Node{
		ID: NodeId,
	}
	node, err := node.ReadByID()
	if err != nil {
		return err
	}
	htmlFiles := []string{
		filepath.Join("./", "templates", "edit", "node.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "node", node)
	return nil
}

func MakerEditPageSubmit(c echo.Context) error {
	NodeId, _ := strconv.Atoi(c.FormValue("id"))
	Text := c.FormValue("text")
	var node models.Node = models.Node{
		ID: NodeId,
	}
	node, err := node.ReadByID()
	if err != nil {
		return err
	}
	node.Text = Text
	if node.ShouldToDo != "" {
		node.ShouldToDo = c.FormValue("should_to_do")
	}
	err = node.Update()
	if err != nil {
		return err
	}
	htmlFiles := []string{
		filepath.Join("./", "templates", "submit", "edit", "node.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "node", node)
	return nil
}
