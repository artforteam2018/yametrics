package app

import (
	"flag"
	"os"
)

type EnvVars struct {
	address        string
	reportInterval string
	pollInterval   string
}

func lookupEnvVal(name string, def *string) {
	if val, ok := os.LookupEnv(name); ok {
		*def = val
	}

}

func Bootstrap() EnvVars {
	envVars := EnvVars{}

	flag.StringVar(&envVars.address, "a", "localhost:8080", "server address to listen on")
	flag.StringVar(&envVars.reportInterval, "r", "10", "report interval")
	flag.StringVar(&envVars.pollInterval, "p", "2", "poll interval")

	flag.Parse()

	lookupEnvVal("ADDRESS", &envVars.address)
	lookupEnvVal("REPORT_INTERVAL", &envVars.reportInterval)
	lookupEnvVal("POLL_INTERVAL", &envVars.pollInterval)

	return envVars
}
