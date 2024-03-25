package routes

import (
	"github.com/artforteam2018/yametrics/internal/server/services/cars"
	"github.com/go-chi/chi/v5"
)

func CarRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", cars.GetCarsHandler)

	r.Route("/{brand}", func(r chi.Router) {
		r.Route("/{model}", func(r chi.Router) {
			r.Get("/", cars.FindBrandAndModelHandler)
		})
		r.Get("/", cars.FindBrandHandler)
	})

	return r
}
