package app

import (
	"time"

	"github.com/artforteam2018/yametrics/internal/agent/components/memstats"
)

func Run() {

	envVars := Bootstrap()
	memstats.Init(envVars.pollInterval, envVars.reportInterval, envVars.address)
	time.Sleep(time.Second * 1000)

}
