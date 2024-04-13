package routes

import (
	"github.com/artforteam2018/yametrics/internal/server/services/metrics"
	"github.com/go-chi/chi/v5"
)

func MetricRoutesUpdate() chi.Router {
	r := chi.NewRouter()

	r.Post("/", metrics.PostMetricJSON)
	r.Route("/{metrictype}/{metricname}/{metricvalue}", func(r chi.Router) {
		r.Post("/", metrics.NoMetric)
	})

	r.Route("/gauge/{metricname}/{metricvalue}", func(r chi.Router) {
		r.Post("/", metrics.GaugeAdd)
	})

	r.Route("/counter/{metricname}/{metricvalue}", func(r chi.Router) {
		r.Post("/", metrics.CounterAdd)
	})

	return r
}

func MetricRoutesGet() chi.Router {
	r := chi.NewRouter()

	r.Route("/gauge/{metricname}", func(r chi.Router) {
		r.Get("/", metrics.GaugeGet)
	})

	r.Route("/counter/{metricname}", func(r chi.Router) {
		r.Get("/", metrics.CounterGet)
	})

	r.Get("/", metrics.GetAllMetrics)
	r.Post("/", metrics.GetMetricJSON)

	return r
}
