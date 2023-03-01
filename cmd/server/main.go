package main

import "fmt"

// App - the struct which contains things like pointer
// to the database connection

type App struct {
}

func (app *App) Run() error {
	fmt.Println("Setting up error app")
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
