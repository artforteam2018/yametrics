package app

import (
	"fmt"
	"net/http"

	"flag"

	"github.com/artforteam2018/yametrics/internal/server/routes"
	"github.com/go-chi/chi/v5"
)

func Run() {

	addrArg := flag.String("a", "localhost:2222", "server address to listen on")
	flag.Parse()

	// re := regexp.MustCompile(`(localhost)|(127.0.0.1)`)
	// address := re.ReplaceAllString(*addrArg, "0.0.0.0")
	address := *addrArg

	r := chi.NewRouter()

	routes.Init(r)

	fmt.Println("server is listening on:", address)
	err := http.ListenAndServe(address, r)

	if err != nil {
		panic(err)
	}
}
