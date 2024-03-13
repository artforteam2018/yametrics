package app

import (
	"fmt"
	"net/http"

	"github.com/artforteam2018/yametrics/internal/server/routes"
)

func Run() {
	mux := http.NewServeMux()

	routes.Init(mux)

	fmt.Println("server is listening on port 8080")
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		panic(err)
	}
}
