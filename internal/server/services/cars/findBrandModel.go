package cars

import (
	"net/http"
	"strings"

	"github.com/artforteam2018/yametrics/internal/server/components/cars"
	"github.com/go-chi/chi/v5"
)

func FindBrandAndModelHandler(w http.ResponseWriter, r *http.Request) {
	brand := chi.URLParam(r, "brand")
	model := chi.URLParam(r, "model")
	cars := cars.FindCarByBrandAndModel(brand, model)

	if len(cars) == 0 {
		http.NotFound(w, r)
		return
	}

	var carModels []string

	for _, car := range cars {
		carModels = append(carModels, car.String())
	}

	w.Write([]byte(strings.Join(carModels, ", ")))

}
