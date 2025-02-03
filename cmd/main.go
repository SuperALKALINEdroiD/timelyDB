package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SuperALKALINEdroiD/timelyDB/config"
	"github.com/SuperALKALINEdroiD/timelyDB/utils/nodes"

	"github.com/go-chi/chi/v5"
)

type App struct {
	Config *config.DatabaseConfig
	Router *chi.Mux
	Nodes  []*nodes.Node
}

func main() {
	config, confifLoadError := initEnvironment()

	if confifLoadError != nil {
		log.Fatalln("error while loading config")
	}

	app := &App{
		Config: config,
		Router: initRouter(config),
		Nodes:  nodes.LoadNodes(config),
	}

	serverAddress := fmt.Sprintf(":%d", app.Config.Port)
	log.Printf("Starting server on %s", serverAddress)

	if err := http.ListenAndServe(serverAddress, app.Router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
