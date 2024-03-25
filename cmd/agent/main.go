package main

import "github.com/artforteam2018/yametrics/internal/agent/app"

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

// 	return strings.Join(out, " ")
// }

func main() {

	// client := Gentleman{}
	// respG := parseRequest(client, url, method)

	// fmt.Printf("response from Gentleman: %v\n", respG)

	// client2 := DefaultRequest{}
	// respD := parseRequest(client2, url, method)

	// fmt.Printf("response from Default client: %v\n", respD)
	app.Run()

}
