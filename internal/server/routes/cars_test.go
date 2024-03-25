package routes

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type RequestInfo struct {
	method string
	body   io.Reader
	uri    string
}

func RequestTest(t *testing.T, h http.Handler, ri RequestInfo) (*http.Response, string) {
	server := httptest.NewServer(h)
	defer server.Close()

	req, err := http.NewRequest(ri.method, server.URL+ri.uri, ri.body)

	require.NoError(t, err)

	resp, err := server.Client().Do(req)

	require.NoError(t, err)

	body, err := io.ReadAll(resp.Body)

	require.NoError(t, err)

	return resp, string(body)
}

func TestCarRoutes(t *testing.T) {
	type want struct {
		StatusCode int
	}

	tests := []struct {
		name string
		req  RequestInfo
		want want
	}{
		{
			name: "brand_model",
			req: RequestInfo{
				method: "GET",
				body:   nil,
				uri:    "/model/brand",
			},
			want: want{
				StatusCode: http.StatusNotFound,
			},
		},
		{
			name: "brand",
			req: RequestInfo{
				method: "GET",
				body:   nil,
				uri:    "/model",
			},
			want: want{
				StatusCode: http.StatusNotFound,
			},
		},
		{
			name: "all",
			req: RequestInfo{
				method: "GET",
				body:   nil,
				uri:    "/",
			},
			want: want{
				StatusCode: http.StatusOK,
			},
		},
		{
			name: "brand_model_found",
			req: RequestInfo{
				method: "GET",
				body:   nil,
				uri:    "/Renault/Logan",
			},
			want: want{
				StatusCode: http.StatusOK,
			},
		},
		{
			name: "brand_found",
			req: RequestInfo{
				method: "GET",
				body:   nil,
				uri:    "/Renault/Logan",
			},
			want: want{
				StatusCode: http.StatusOK,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			resp, _ := RequestTest(t, CarRoutes(), test.req)
			defer resp.Body.Close()
			assert.Equal(tt, test.want.StatusCode, resp.StatusCode)
		})
	}
}
