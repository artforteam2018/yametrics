package app

import (
	"time"

	"github.com/artforteam2018/yametrics/internal/agent/components/memstats"
)

func Run() {
	memstats.Init()
	time.Sleep(1000 * time.Minute)
}
