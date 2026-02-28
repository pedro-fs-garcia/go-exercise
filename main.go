package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
	"training/config"
	_ "training/config"
	"training/internal/app"
	"training/internal/routes"
)

func main() {
	err := config.LoadEnvironment()
	if err != nil {
		panic(err)
	}

	var port int
	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	defer app.DB.Close()

	app.Logger.Printf("Application running on port %d...\n", port)

	r := routes.SetupRoutes(app)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
