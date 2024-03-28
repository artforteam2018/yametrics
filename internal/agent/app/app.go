package app

import (
	"flag"
	"time"

	"github.com/artforteam2018/yametrics/internal/agent/components/memstats"
)

func Run() {
	addrArg := flag.String("a", "localhost:8080", "server address to listen on")
	flag.Parse()

	address := *addrArg

	reportInterval := flag.Int("r", 10, "report interval")
	pollInterval := flag.Int("p", 2, "poll interval")

	flag.Parse()

	memstats.Init(*pollInterval, *reportInterval, address)
	time.Sleep(time.Second * 1000)

}
