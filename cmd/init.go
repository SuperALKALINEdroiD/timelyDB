package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SuperALKALINEdroiD/timelyDB/config"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func initEnvironment() (*config.DatabaseConfig, error) {
	cfg, err := config.LoadConfig("") // TODO: Load from environment or file path
	if err != nil {
		log.Printf("Error loading configuration: %v", err)
		return nil, err
	}
	return cfg, nil
}

func initRouter(cfg *config.DatabaseConfig) *chi.Mux {
	router := chi.NewRouter()
	addMiddlewares(router)
	initRoutes(router, cfg)
	return router
}

func addMiddlewares(router *chi.Mux) {
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
}

func initRoutes(router *chi.Mux, cfg *config.DatabaseConfig) {
	router.Route("/data-in", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "Server is running")
		})
		r.Post("/upsert", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Upsert Endpoint WIP - Config: %+v", cfg)
		})
		r.Post("/insert", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Insert Endpoint WIP - Config: %+v", cfg)
		})
		r.Post("/update", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Update Endpoint WIP - Config: %+v", cfg)
		})
	})
}
