package main

import (
	"fmt"
	"io"
	"net/http"
)

type DefaultRequest struct {
}

func (r DefaultRequest) MakeRequest(url string, method string) ([]byte, error) {
	client := http.DefaultClient
	req, err := http.NewRequest(method, url+"/users", nil)

	if err != nil {
		return nil, fmt.Errorf("failed to build request: %v", err)
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("failed to read request: %v", err)
	}

	return body, nil
}
