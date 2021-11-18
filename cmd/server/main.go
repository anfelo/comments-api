package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/anfelo/comments-api/internal/transport/http"
)

// App - the struct which contains things like pointers
// to database connections
type App struct{}

// Run - sets up our application
func (a *App) Run() error {
	fmt.Println("Setting Up Out APP")

	handler := transportHTTP.NewHandler()
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
