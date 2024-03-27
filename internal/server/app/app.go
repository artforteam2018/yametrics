package app

import (
	"fmt"
	"net/http"
	"regexp"

	"flag"

	"github.com/artforteam2018/yametrics/internal/server/routes"
	"github.com/go-chi/chi/v5"
)

func Run() {

	addrArg := flag.String("a", "localhost:8080", "server address to listen on")
	flag.Parse()

	re := regexp.MustCompile(`(localhost)|(127.0.0.1)`)
	address := re.ReplaceAllString(*addrArg, "")
	// address := *addrArg

	r := chi.NewRouter()

	routes.Init(r)

	fmt.Println("server is listening on:", address)
	err := http.ListenAndServe(address, r)

	if err != nil {
		panic(err)
	}
}
