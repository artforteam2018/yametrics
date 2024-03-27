package app

import (
	"flag"
	"time"

	"github.com/artforteam2018/yametrics/internal/agent/components/memstats"
)

func Run() {
	reportInterval := flag.Int("r", 10, "report interval")
	pollInterval := flag.Int("p", 2, "poll interval")

	flag.Parse()

	memstats.Init(*pollInterval, *reportInterval)
	time.Sleep(1000 * time.Minute)
}
