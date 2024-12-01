package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func initEnvironment() {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading Environment")
  }
}

func initRouter() *chi.Mux {
		router := chi.NewRouter()
		addMiddlewares(router)
		initRoutes(router)
		return router
}

func addMiddlewares(r *chi.Mux)  {
	r.Use(middleware.Logger)
}

func initRoutes(router *chi.Mux) {
	router.Route("/data-in", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.Context().Value("DatabaseConfig"))
		})
	})
}
