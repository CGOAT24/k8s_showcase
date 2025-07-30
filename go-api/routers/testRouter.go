package routers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type TestRessource struct{}

func (rs TestRessource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.Get)

	return r
}

func (rs TestRessource) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a test"))
}