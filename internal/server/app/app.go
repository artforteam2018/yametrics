package app

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/artforteam2018/yametrics/internal/server/routes"
)

func Run() {

	r := chi.NewRouter()

	routes.Init(r)

	fmt.Println("server is listening on port 8080")
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		panic(err)
	}
}
