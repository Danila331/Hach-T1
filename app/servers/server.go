package servers

import (
	"github.com/Danila331/HACH-T1/app/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer() error {
	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	files := app.Group("/files")
	files.GET("/add", handlers.AddFilePage)
	files.POST("/add/submit", handlers.AddFileSubmit)
	return app.Start(":8080")
}
