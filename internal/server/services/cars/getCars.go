package cars

import (
	"net/http"
	"strings"

	"github.com/artforteam2018/yametrics/internal/server/components/cars"
)

func GetCarsHandler(w http.ResponseWriter, r *http.Request) {
	cars := cars.GetCarModels()

	w.Write([]byte(strings.Join(cars, ",")))

}
