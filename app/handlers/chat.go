package handlers

import (
	"path/filepath"
	"text/template"
	"time"

	"github.com/Danila331/HACH-T1/app/pkg"
	"github.com/labstack/echo/v4"
)

func ChatPage(c echo.Context) error {
	time.Sleep(5 * time.Second)
	htmlFiles := []string{filepath.Join("./", "templates", "chat.html")}
	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}
	templ.ExecuteTemplate(c.Response(), "chat", nil)
	return nil
}

func SendMessage(c echo.Context) error {
	message := c.FormValue("message")
	BotMessage, err := pkg.SendMessage("http://192.168.250.93:5002/api/llm/anthropic/SONNET3_5", message)
	if err != nil {
		return err
	}
	htmlFiles := []string{filepath.Join("./", "templates", "submit", "chat_submit.html")}
	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}
	templ.ExecuteTemplate(c.Response(), "chat_submit", BotMessage)
	return nil
}
