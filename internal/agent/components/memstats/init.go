package memstats

import (
	"github.com/artforteam2018/yametrics/internal/agent/constants"
	"github.com/artforteam2018/yametrics/internal/agent/services/interval"
)

func Init() {
	mstats := MemStats{}
	mstats.Init()
	interval.StartInterval(constants.PollInterval, func() {
		mstats.Scan()
	})

	interval.StartInterval(constants.ReportInterval, func() {
		mstats.Send()
	})
}
