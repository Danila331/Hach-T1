package handlers

import (
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/Danila331/HACH-T1/app/models"
	"github.com/labstack/echo/v4"
)

func AddDatabasePage(c echo.Context) error {
	return c.String(http.StatusOK, "database page")
}

func AddPostgressqlPage(c echo.Context) error {
	htmlFiles := []string{
		filepath.Join("./", "templates", "add_postgres.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "add_postgres", nil)
	return nil
}

// Функция AddPostgressqlSubmit принимает данные из формы и создает новую запись в базе данных
func AddPostgressqlSubmit(c echo.Context) error {
	host := c.FormValue("host")
	port := c.FormValue("port")
	user := c.FormValue("user")
	password := c.FormValue("password")
	dbname := c.FormValue("dbname")
	var db = models.Database{
		Type:         "Postgressql",
		Host:         host,
		Port:         port,
		UserName:     user,
		Password:     password,
		DatabaseName: dbname,
	}
	err := db.Create()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "Postgressql submit")
}

func AddMysqlPage(c echo.Context) error {
	htmlFiles := []string{
		filepath.Join("./", "templates", "add_mysql.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "add_mysql", nil)
	return nil
}

// Функция AddMysqlSubmit принимает данные из формы и создает новую запись в базе данных
func AddMysqlSubmit(c echo.Context) error {
	host := c.FormValue("host")
	port := c.FormValue("port")
	user := c.FormValue("user")
	password := c.FormValue("password")
	dbname := c.FormValue("dbname")
	var db = models.Database{
		Type:         "Mysql",
		Host:         host,
		Port:         port,
		UserName:     user,
		Password:     password,
		DatabaseName: dbname,
	}
	err := db.Create()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "Mysql submit")
}

func AddElasticsearchPage(c echo.Context) error {
	htmlFiles := []string{
		filepath.Join("./", "templates", "add_elastic.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "add_elastic", nil)
	return nil
}

// Функция AddElasticsearchSubmit принимает данные из формы и создает новую запись в базе данных
func AddElasticsearchSubmit(c echo.Context) error {
	host := c.FormValue("host")
	port := c.FormValue("port")
	user := c.FormValue("user")
	password := c.FormValue("password")
	var db = models.Database{
		Type:         "Elasticsearch",
		Host:         host,
		Port:         port,
		UserName:     user,
		Password:     password,
		DatabaseName: "",
	}
	err := db.Create()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "Elasticsearch submit")
}

func AddMongodbPage(c echo.Context) error {
	htmlFiles := []string{
		filepath.Join("./", "templates", "add_mongo.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "add_mongo", nil)
	return nil
}

// Функция AddMongodbSubmit принимает данные из формы и создает новую запись в базе данных
func AddMongodbSubmit(c echo.Context) error {
	host := c.FormValue("host")
	port := c.FormValue("port")
	user := c.FormValue("user")
	password := c.FormValue("password")
	databasename := c.FormValue("dbname")
	var db = models.Database{
		Type:         "Mongodb",
		Host:         host,
		Port:         port,
		UserName:     user,
		Password:     password,
		DatabaseName: databasename,
	}
	err := db.Create()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "Mongodb submit")
}
