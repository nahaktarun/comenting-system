package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/nahaktarun/go-rest-api/internal/transport/http"
)

// App - the struct which contains things like pointer
// to the database connection

type App struct {
}

func (app *App) Run() error {
	fmt.Println("Setting up error app")
	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to setup server")
		return err
	}
	return nil
}
func main() {
	fmt.Println("Go rest api")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up rest api ")
		fmt.Println(err)
	}
}
