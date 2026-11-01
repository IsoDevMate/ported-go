package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
}

type config struct {
	addr string
	env  string
	db   dbConfig
}

type dbConfig struct {
	dbURL        string
	authToken    string
	maxOpenConns int
	maxIdleConns int
}

// Chi handles your primary application routing (taking advantage of Chi-specific middleware),
// but passes a specific sub-route off to the standard library http.ServeMux.
func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	//r.Use(middleware.RealIP) this library is deprecated and may expose application to ip spoofing
	r.Use(middleware.Recoverer)

	r.Get("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})
	r.Get("/health", app.healthHandler)

	return r
}

func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("listening on %s", srv.Addr)
	return srv.ListenAndServe()
}
