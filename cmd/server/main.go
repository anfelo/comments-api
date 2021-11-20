package main

import (
	"fmt"
	"net/http"

	"github.com/anfelo/comments-api/internal/comment"
	"github.com/anfelo/comments-api/internal/database"
	transportHTTP "github.com/anfelo/comments-api/internal/transport/http"
)

// App - the struct which contains things like pointers
// to database connections
type App struct{}

// Run - sets up our application
func (a *App) Run() error {
	fmt.Println("Setting Up Out APP")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go Rest API Course")
	app := App{}

	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
