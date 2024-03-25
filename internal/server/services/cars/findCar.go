package cars

import (
	"net/http"

	"github.com/artforteam2018/yametrics/internal/server/components/cars"
	"github.com/go-chi/chi/v5"
)

func FindCarHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	car, err := cars.FindCarByID(id)
	if err != nil {
		http.NotFound(w, r)
	}
	w.Write([]byte(car.String()))

}
