package app

import (
	"fmt"
	"net/http"

	"github.com/artforteam2018/yametrics/internal/server/routes"
	"github.com/go-chi/chi/v5"
)

func colorifyRed(s string) string {
	return "\033[1;31m" + s + "\033[0m"
}

func colorifyGreen(s string) string {
	return "\033[1;32m" + s + "\033[0m"
}

func Run() {

	envVars := Bootstrap()

	r := chi.NewRouter()

	routes.Init(r)

	fmt.Println("Welcome, " + colorifyGreen(envVars.user) + ", server is listening on " + colorifyRed(envVars.address))
	err := http.ListenAndServe(envVars.address, r)

	if err != nil {
		panic(err)
	}
}
