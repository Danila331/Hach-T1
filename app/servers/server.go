package servers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func StartServer() error {
	app := echo.New()
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	return app.Start(":8080")
}
