package main

import (
	"frontend/pkg/config"
	"frontend/pkg/handlers"
	"frontend/pkg/render"
	"frontend/pkg/web"
	"log"
)

const webPort = "80"

func main() {
	// app holds our application configuration
	var app config.AppConfig

	// Customize log display
	setLog()

	// Create Template Cache
	// will be created at compilation time
	// to have it rebuilt for each handler request (during runtime)
	// set app.UseCache to false (dev environment)
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// set the config for the template package
	app.TemplateCache = tc
	app.UseCache = false // Must be set to true for production
	render.NewTemplate(&app)

	// set the config for the handlers package
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// Start web server
	if err := web.Run(multiplexer(), webPort); err != nil {
		log.Fatalf(err.Error())
	}
}
