package app

import (
	"flag"
	"io"
	"log"
	"os"

	Logger "github.com/artforteam2018/yametrics/internal/server/components/logger"
)

type EnvVars struct {
	address string
	user    string
}

func (e EnvVars) String() string {
	return "address=" + e.address + "user=" + e.user
}

func lookupEnvVal(name string, valLink *string) {
	if val, ok := os.LookupEnv(name); ok {
		*valLink = val
	}

}

type Duplex struct {
	streams []io.Writer
}

func (d Duplex) Write(p []byte) (n int, err error) {
	for _, stream := range d.streams {
		nDone, errWrite := stream.Write(p)
		if err != nil {
			err = errWrite
		}
		n = nDone
	}
	return n, err
}

func Bootstrap() EnvVars {
	envVars := EnvVars{}
	duplex := Duplex{}

	flag.StringVar(&envVars.address, "a", "localhost:8080", "server address to listen on")

	flag.Parse()

	lookupEnvVal("ADDRESS", &envVars.address)
	lookupEnvVal("USER", &envVars.user)

	file, err := os.OpenFile("./info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		panic("cannot open log file" + err.Error())
	}

	duplex.streams = []io.Writer{file, log.Writer()}

	log.SetOutput(duplex)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.Println("Successfully started with env params:", envVars)

	Logger.Init()
	return envVars
}
