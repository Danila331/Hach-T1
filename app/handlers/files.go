package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/Danila331/HACH-T1/app/models"
	"github.com/Danila331/HACH-T1/app/pkg"
	"github.com/labstack/echo/v4"
)

func AddFilePage(c echo.Context) error {
	htmlFiles := []string{
		filepath.Join("./", "templates", "file_add.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "file_add", nil)
	return nil
}

func AddFileSubmit(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.String(http.StatusBadRequest, "Failed to parse form")
	}

	files := form.File["files[]"]
	if len(files) == 0 {
		return c.String(http.StatusBadRequest, "No files uploaded")
	}

	saveDir := "uploads"
	if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to create upload directory")
	}

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to open file: %s", file.Filename))
		}
		defer src.Close()

		// Destination file path
		destPath := filepath.Join(saveDir, file.Filename)
		dest, err := os.Create(destPath)
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to create file: %s", file.Filename))
		}
		defer dest.Close()

		// Copy file content
		if _, err := dest.ReadFrom(src); err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to save file: %s", file.Filename))
		}
		// Закончился код загрузки файла
		var modelFile = models.File{
			Name: file.Filename,
			Path: string(destPath),
		}

		err = modelFile.Create()
		if err != nil {
			return err
		}

		err = pkg.S3LoadFile(file.Filename, destPath)
		if err != nil {
			return err
		}
	}

	println("Files uploaded successfully")

	htmlFiles := []string{
		filepath.Join("./", "templates", "submit", "file_submit.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "file_submit", nil)
	return nil
}
