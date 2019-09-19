package main

import (
	"net/http"

	"github.com/spankie/go-web-mysql/db"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/spankie/go-web-mysql/handlers"
)

func main() {
	r := chi.NewRouter() // creating a new router
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// get the pointer to the DB so that it can be closed
	// after the main function returns
	DB := db.GetDB()
	defer DB.Close() // Idiomatic go

	fileServer := http.StripPrefix("/assets/", http.FileServer(http.Dir("./public/assets")))
	r.Get("/assets/*", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Vary", "Accept-Encoding")
		w.Header().Set("Cache-Control", "public, max-age=7776000")
		fileServer.ServeHTTP(w, r)
	})
	//PARSE
	r.Get("/", handlers.HomeHandler)
	r.Post("/signup", handlers.SignupHandler)
	r.Get("/login", handlers.SignupHandler)
	r.Post("/login", handlers.SignupHandler)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/404.html")
	})
	log.Print("👉  Server started at 127.0.0.1:8080")
	http.ListenAndServe(":8080", r)
}
