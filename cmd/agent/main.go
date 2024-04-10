package main

import (
	"fmt"
	"time"

	"github.com/artforteam2018/yametrics/internal/agent/app"
)

// "encoding/json"
// "fmt"
// "strings"

// type User struct {
// 	ID       int    `json:"id"`
// 	Username string `json:"username"`
// 	Email    string `json:"email"`
// }

// const url = "https://jsonplaceholder.typicode.com"
// const method = "GET"

// type RequestClient interface {
// 	MakeRequest(url string, method string) ([]byte, error)
// }

// func parseRequest(client RequestClient, url string, method string) string {
// 	body, err := client.MakeRequest(url, method)
// 	var resp []User

// 	if err != nil {
// 		fmt.Printf("Error requesting: %v\n", err)
// 	} else {
// 		json.Unmarshal(body, &resp)
// 	}

// 	var out []string

// 	for _, u := range resp {
// 		out = append(out, u.Username)
// 	}

//		return strings.Join(out, " ")
//	}
func toFixed(num float64, precision int) float64 {
	output := 1.0
	for i := 0; i < precision; i++ {
		output *= 10
	}
	return float64(int(num*output)) / output
}

func main() {
	timer := time.Now()
	for {
		fmt.Println(int(time.Since(timer).Milliseconds()))
		time.Sleep(time.Millisecond * 10)
	}
	// client := Gentleman{}
	// respG := parseRequest(client, url, method)

	// fmt.Printf("response from Gentleman: %v\n", respG)

	// client2 := DefaultRequest{}
	// respD := parseRequest(client2, url, method)

	// fmt.Printf("response from Default client: %v\n", respD)
	app.Run()

}
