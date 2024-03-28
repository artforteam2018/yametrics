package app

import (
	"flag"
	"os"
)

type EnvVars struct {
	address string
	user    string
}

func lookupEnvVal(name string, valLink *string) {
	if val, ok := os.LookupEnv(name); ok {
		*valLink = val
	}

}

func Bootstrap() EnvVars {
	envVars := EnvVars{}

	flag.StringVar(&envVars.address, "a", "localhost:8080", "server address to listen on")

	flag.Parse()

	lookupEnvVal("ADDRESS", &envVars.address)
	lookupEnvVal("USER", &envVars.user)

	return envVars
}
