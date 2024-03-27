package app

import (
	"flag"
	"fmt"
	"net/http"
	"regexp"

	"github.com/artforteam2018/yametrics/internal/agent/components/memstats"
)

func Run() {
	addrArg := flag.String("a", "localhost:8080", "server address to listen on")

	reportInterval := flag.Int("r", 10, "report interval")
	pollInterval := flag.Int("p", 2, "poll interval")

	flag.Parse()

	re := regexp.MustCompile(`(localhost)|(127.0.0.1)`)
	address := re.ReplaceAllString(*addrArg, "")

	memstats.Init(*pollInterval, *reportInterval)

	fmt.Println("server is listening on:", address)
	err := http.ListenAndServe(address, nil)

	if err != nil {
		panic(err)
	}
}
