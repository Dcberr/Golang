package main

import (
	"context"
	"github/Dcberr/dto"
	"github/Dcberr/templates"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Render the template using Go's built-in html/template
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    t, err := template.ParseFiles(tmpl)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    t.Execute(w, data)
}

func main() {
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    cards := []*dto.Card{
        {Title: "Computer Network Files", Participants: "24", ImgUrl: "https://via.placeholder.com/80", ButtonText: "Join and share file now!"},
        {Title: "Database Files", Participants: "60", ImgUrl: "https://via.placeholder.com/80", ButtonText: "Join and share file now!"},
    }

    e.GET("/", func(c echo.Context) error {
  		index := templates.Index(cards)
        return index.Render(context.Background(), c.Response().Writer)
    })

    e.Logger.Fatal(e.Start(":8080"))
}
