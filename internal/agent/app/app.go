// package app

// import (
// 	"context"
// 	"fmt"
// 	"net"
// 	"net/http"
// 	"time"

// 	"github.com/artforteam2018/yametrics/internal/agent/components/memstats"
// 	"github.com/go-chi/chi"
// )

// func Run() {
// 	memstats.Init()
// 	r := chi.NewRouter()

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
// 		defer cancel()
// 		lc := &net.ListenConfig{}
// 		fmt.Println("start listen")
// 		ln, _ := lc.Listen(ctx, "tcp", ":8080")
// 		fmt.Println(ln)
// 		if ln != nil {
// 			defer ln.Close()
// 			done := make(chan struct{})
// 			go func() {
// 				fmt.Println("go esende")
// 				conn, _ := ln.Accept()
// 				fmt.Println("go esende2", conn)

// 				if conn != nil {
// 					_ = conn.Close()
// 				}
// 				close(done)
// 			}()
// 			select {
// 			case <-done:
// 				fmt.Println("done listen")
// 				return
// 			case <-ctx.Done():
// 				fmt.Println("no listen")
// 				fmt.Println(ctx.Err())
// 			}
// 		}
// 	}()
// 	fmt.Println("run")
// 	err := http.ListenAndServe(`localhost:8080`, r)

// 	if err != nil {
// 		panic(err)
// 	}

// }

package app

import (
	"flag"
	"fmt"
	"net/http"
	"regexp"

	"github.com/artforteam2018/yametrics/internal/agent/components/memstats"
)

func Run() {
	addrArg := flag.String("a", "localhost:8080", "server address to listen on")

	reportInterval := flag.Int("r", 10, "report interval")
	pollInterval := flag.Int("p", 2, "poll interval")

	flag.Parse()

	re := regexp.MustCompile(`(localhost)|(127.0.0.1)`)
	address := re.ReplaceAllString(*addrArg, "")

	memstats.Init(*pollInterval, *reportInterval)

	fmt.Println("server is listening on:", address)
	err := http.ListenAndServe(address, nil)

	if err != nil {
		panic(err)
	}
}
