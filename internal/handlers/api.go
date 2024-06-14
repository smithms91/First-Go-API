package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/smithms91/goapi/internal/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(r chi.Router) {
		r.Use(middleware.Authorization)
		r.Get("/coins", GetCoinBalance)
	})
}
