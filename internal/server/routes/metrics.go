package routes

import (
	"github.com/artforteam2018/yametrics/internal/server/services/metrics"
	"github.com/go-chi/chi/v5"
)

func MetricRoutesUpdate() chi.Router {
	r := chi.NewRouter()

	r.Post("/", metrics.HandleNoMetric)
	r.Route("/{metrictype}/{metricname}/{metricvalue}", func(r chi.Router) {
		r.Post("/", metrics.HandleNoMetric)
	})

	r.Route("/gauge/{metricname}/{metricvalue}", func(r chi.Router) {
		r.Post("/", metrics.HandleGaugeAdd)
	})

	r.Route("/counter/{metricname}/{metricvalue}", func(r chi.Router) {
		r.Post("/", metrics.HandleCounterAdd)
	})

	return r
}

func MetricRoutesGet() chi.Router {
	r := chi.NewRouter()

	r.Route("/gauge/{metricname}", func(r chi.Router) {
		r.Get("/", metrics.HandleGaugeGet)
	})

	r.Route("/counter/{metricname}", func(r chi.Router) {
		r.Get("/", metrics.HandleCounterGet)
	})

	r.Get("/", metrics.HandleAllMetrics)

	return r
}
