package memstats

import (
	"strconv"
	"time"

	"github.com/artforteam2018/yametrics/internal/agent/services/interval"
)

func Init(pollinterval string, reportinterval string, address string) {
	pollInterval, err := strconv.Atoi(pollinterval)
	reportInterval, err1 := strconv.Atoi(reportinterval)

	if err != nil || err1 != nil {
		panic("Env variables cannot be parsed to Int" + pollinterval + reportinterval)
	}

	mstats := InitMemstats(address)
	mstats.Init()
	interval.StartInterval(time.Second*time.Duration(pollInterval), func() {
		mstats.Scan()
	})

	interval.StartInterval(time.Second*time.Duration(reportInterval), func() {
		mstats.Send()
	})
}
