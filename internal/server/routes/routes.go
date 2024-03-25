package routes

import (
	"fmt"
	"io"
	"net/http"

	middlewareC "github.com/artforteam2018/yametrics/internal/server/middlewares"
	"github.com/artforteam2018/yametrics/internal/server/services/cars"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Init(r *chi.Mux) {

	r.Use(middleware.Recoverer)
	r.Use(middlewareC.TimerPrint)

	r.Mount("/update", MetricRoutesUpdate())
	r.Mount("/value", MetricRoutesGet())

	r.Mount("/cars", CarRoutes())

	r.Route("/car/{id}", func(r chi.Router) {
		r.Get("/", cars.FindCarHandler)
	})

	r.Post("/test", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)

		if err != nil {
			fmt.Println(err.Error())
		}

		for k, v := range r.Header {
			fmt.Println("HEADER", k, v)
		}
		fmt.Println("URL", r.URL)
		fmt.Println("body", string(body))
	})

}
