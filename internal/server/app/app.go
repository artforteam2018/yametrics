package app

import (
	"fmt"
	"net/http"

	"flag"

	"github.com/artforteam2018/yametrics/internal/server/routes"
	"github.com/go-chi/chi/v5"
)

func Run() {

	addr := flag.String("a", "localhost:8080", "server address to listen on")

	flag.Parse()

	r := chi.NewRouter()

	routes.Init(r)

	fmt.Println("server is listening on port 8080")
	err := http.ListenAndServe(*addr, r)

	if err != nil {
		panic(err)
	}
}
