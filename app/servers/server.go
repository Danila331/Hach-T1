package servers

import (
	"github.com/Danila331/HACH-T1/app/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer() error {
	app := echo.New()
	app.HTTPErrorHandler = handlers.CustomHTTPErrorHandler
	app.Static("/static", "./static")
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/chat", handlers.ChatPage)
	app.POST("/send-message", handlers.SendMessage)
	// Group for maker
	maker := app.Group("/maker")
	maker.GET("/", handlers.MakerPage)
	maker.GET("/add", handlers.MakerAddPage)
	// Group for maker add triger
	triger := maker.Group("/add/triger")
	triger.GET("/", handlers.MakerTrigerPage)
	triger.POST("/submit", handlers.MakerTrigerPageSubmit)
	// Group for maker add start
	start := maker.Group("/add/start")
	start.GET("/", handlers.MakerStartPage)
	start.POST("/submit", handlers.MakerStartPageSubmit)
	// Group for maker add end
	end := maker.Group("/add/end")
	end.GET("/", handlers.MakerEndPage)
	end.POST("/submit", handlers.MakerEndPageSubmit)
	//Group for maker add error
	er := maker.Group("/add/error")
	er.GET("/", handlers.MakerErrorPage)
	er.POST("/submit", handlers.MakerErrorPageSubmit)
	// Group for maker edit
	editMaker := maker.Group("/edit")
	editMaker.GET("/", handlers.MakerEditPage)
	editMaker.POST("/submit", handlers.MakerEditPageSubmit)
	// Group for maker delete
	deleteMaker := maker.Group("/delete")
	deleteMaker.POST("/submit", handlers.MakerDeleteSubmit)

	// Group for uploads
	uploads := app.Group("/uploads")
	uploads.GET("/", handlers.UploadsPage)
	uploads.GET("/add", handlers.UploadsAddPage)
	// Group for edit
	edit := uploads.Group("/edit")
	edit.GET("/database", handlers.EditDataBasePage)
	edit.POST("/database/submit", handlers.EditDataBasePageSubmit)
	edit.GET("/website", handlers.EditWebsitePage)
	edit.POST("/website/submit", handlers.EditWebsitePageSubmit)
	// Group for delete
	delete := uploads.Group("/delete")
	delete.POST("/file/submit", handlers.DeleteFileSubmit)
	delete.POST("/database/submit", handlers.DeleteDatabaseSubmit)
	delete.POST("/website/submit", handlers.DeleteWebsiteSubmit)
	// Group for files
	file := uploads.Group("/add/file")
	file.GET("/", handlers.AddFilePage)
	file.POST("/submit", handlers.AddFileSubmit)
	// Group for database
	database := uploads.Group("/add/database")
	database.GET("/", handlers.AddDatabasePage)
	// Group for postgressql
	postgressql := database.Group("/postgressql")
	postgressql.GET("/", handlers.AddPostgressqlPage)
	postgressql.POST("/submit", handlers.AddPostgressqlSubmit)
	// Group for mysql
	mysql := database.Group("/mysql")
	mysql.GET("/", handlers.AddMysqlPage)
	mysql.POST("/submit", handlers.AddMysqlSubmit)
	// Group for elasticsearch
	elasticsearch := database.Group("/elasticsearch")
	elasticsearch.GET("/", handlers.AddElasticsearchPage)
	elasticsearch.POST("/submit", handlers.AddElasticsearchSubmit)
	// Group for mongodb
	mongodb := database.Group("/mongodb")
	mongodb.GET("/", handlers.AddMongodbPage)
	mongodb.POST("/submit", handlers.AddMongodbSubmit)
	// Group for website
	website := uploads.Group("/add/website")
	website.GET("/", handlers.AddWebsitePage)
	website.POST("/submit", handlers.AddWebsiteSubmit)
	return app.Start("172.16.106.89:8081")
}
