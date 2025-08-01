package main

import (
	"github.com/go-chi/chi/v5/middleware"
	"net/http"

	"api/routers"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world from Go!"))
	})

	r.Mount("/test", routers.TestRessource{}.Routes())

	http.ListenAndServe(":8000", r)
}
