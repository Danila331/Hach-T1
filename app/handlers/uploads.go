package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func UploadsPage(c echo.Context) error {
	return c.String(http.StatusOK, "Uploads page")
}

func UploadsAddPage(c echo.Context) error {
	return c.String(http.StatusOK, "Uploads handler")
}
