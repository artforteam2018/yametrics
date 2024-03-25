package main

import (
	"fmt"

	"gopkg.in/h2non/gentleman.v2"
)

type Gentleman struct {
}

func (g Gentleman) MakeRequest(url string, method string) (body []byte, err error) {

	cli := gentleman.New()
	cli.URL(url)

	req := cli.Request()
	req.Path("/users")
	req.Method(method)

	res, err := req.Send()

	if err != nil {
		return nil, fmt.Errorf("error sending request, %v", err)
	}

	if !res.Ok {
		return nil, fmt.Errorf("invalid server response, %d", res.StatusCode)
	}

	return res.Bytes(), nil
}
