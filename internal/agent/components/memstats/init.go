package memstats

import (
	"time"

	"github.com/artforteam2018/yametrics/internal/agent/services/interval"
)

func Init(pollinterval int, reportinterval int, address string) {
	mstats := InitMemstats(address)
	mstats.Init()
	interval.StartInterval(time.Second*time.Duration(pollinterval), func() {
		mstats.Scan()
	})

	interval.StartInterval(time.Second*time.Duration(reportinterval), func() {
		mstats.Send()
	})
}
