// package main

// import (
// 	"context"
// 	// "database/sql"

// 	"io"
// 	"log"
// 	"os"
// 	"path/filepath"

// 	// "github/Dcberr/dto"
// 	"github/Dcberr/templates"
// 	"html/template"
// 	"net/http"

// 	"github.com/jmoiron/sqlx"
// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"
// 	_ "github.com/mattn/go-sqlite3"
// )

// // Render the template using Go's built-in html/template
// func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
//     t, err := template.ParseFiles(tmpl)
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }
//     t.Execute(w, data)
// }
// func main() {
//     e := echo.New()

//     //Database Setup
//     // Set up a database connection (using SQLite for simplicity)
// 	db, err := sqlx.Open("sqlite3", "./files.db")
// 	if err != nil {
// 		e.Logger.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Initialize the database table
// 	initDB(db)

//     	// Define the file upload route
// 	e.POST("/upload", func(c echo.Context) error {
// 		// Source
// 		file, err := c.FormFile("file")
// 		if err != nil {
// 			return c.String(http.StatusBadRequest, "File is required")
// 		}

//         //Opening and Saving the file
//         src, err := file.Open()
// 		if err != nil {
// 			return err
// 		}
// 		defer src.Close()

// 		// Create destination file
// 		filename := filepath.Base(file.Filename)
// 		dstPath := "./uploads/" + filename

// 		dst, err := os.Create(dstPath)
// 		if err != nil {
// 			return err
// 		}
// 		defer dst.Close()

// 		// Copy file contents
// 		if _, err = io.Copy(dst, src); err != nil {
// 			return err
// 		}

//         // Save file info to the database
// 		if err := saveFileInfo(db, dstPath); err != nil {
// 			return err
// 		}

// 		return c.String(http.StatusOK, "Done")

// 	})

//             e.GET("/", func(c echo.Context) error {
//                 index := templates.Index3()
//             return index.Render(context.Background(), c.Response().Writer)
//         })

//         // e.GET("/", func(c echo.Context) error {
//         //     return c.File("index.html")
//         // })

//     e.Use(middleware.Logger())
//     e.Use(middleware.Recover())

//     e.Logger.Fatal(e.Start(":8080"))

// }

// func initDB(db *sqlx.DB) {
//     query := `CREATE TABLE IF NOT EXISTS files (
//         id INTEGER PRIMARY KEY AUTOINCREMENT,
//         file_path TEXT NOT NULL,
//         uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
//     );`
//     _, err := db.Exec(query)
//     if err != nil {
//         log.Fatal(err)
//     }
// }

// func saveFileInfo(db *sqlx.DB, filePath string) error {
//     query := `INSERT INTO files (file_path) VALUES (?)`
//     _, err := db.Exec(query, filePath)
//     return err
// }

package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TemplateRenderer struct {
    templates *template.Template
}

func NewTemplateRenderer() *TemplateRenderer {
    return &TemplateRenderer{
        templates: template.Must(template.ParseGlob("templates/*.html")),
    }
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
	e.Renderer = NewTemplateRenderer()
    
    // Serve static files
    e.Static("/static", "static")
    
    // Home page route
    e.GET("/", func(c echo.Context) error {
        return c.Render(http.StatusOK, "home.html", nil)
    })
    
    // Channel page route
    e.GET("/channel/:name", func(c echo.Context) error {
        channelName := c.Param("name")
        data := map[string]interface{}{
            "ChannelName": channelName,
        }
        return c.Render(http.StatusOK, "channel.html", data)
    })

	e.GET("/new-channel-form", func(c echo.Context) error {
		return c.Render(http.StatusOK, "new_channel_form.html", nil)
	})
		
	e.POST("/create-channel", func(c echo.Context) error {
		channelName := c.FormValue("channel_name")
		ownerName := c.FormValue("owner_name")
	
		// Render channel box template with the provided data
		data := map[string]interface{}{
			"ChannelName": channelName,
			"OwnerName":   ownerName,
		}
		return c.Render(http.StatusOK, "channel_box.html", data)
	})

	
    
    // Start server
    e.Logger.Fatal(e.Start(":8080"))
}







