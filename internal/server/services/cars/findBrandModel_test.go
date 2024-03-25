package cars

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type RequestInfo struct {
	method string
	body   io.Reader
	params map[string]string
}

func RequestTest(t *testing.T, h http.Handler, ri RequestInfo) (*http.Response, string) {

	req, err := http.NewRequest(ri.method, "/", ri.body)
	require.NoError(t, err)

	rctx := chi.NewRouteContext()

	for key, val := range ri.params {
		rctx.URLParams.Add(key, val)
	}

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	resp := httptest.NewRecorder()
	h.ServeHTTP(resp, req)

	body, err := io.ReadAll(resp.Body)

	require.NoError(t, err)

	return resp.Result(), string(body)
}

func TestFindBrandAndModelHandler(t *testing.T) {
	type want struct {
		StatusCode int
		body       string
	}

	tests := []struct {
		name string
		req  RequestInfo
		want want
	}{
		{
			name: "ok",
			req: RequestInfo{
				method: "GET",
				body:   nil,
				params: map[string]string{
					"brand": "Renault",
					"model": "Logan",
				},
			},
			want: want{
				StatusCode: 200,
				body:       "Brand: Renault, Model: Logan",
			},
		},
		{
			name: "not found",
			req: RequestInfo{
				method: "GET",
				body:   nil,
				params: map[string]string{
					"brand": "Not",
					"model": "Found",
				},
			},
			want: want{
				StatusCode: 404,
				body:       "",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			resp, body := RequestTest(t, http.HandlerFunc(FindBrandAndModelHandler), test.req)
			assert.Equal(tt, test.want.StatusCode, resp.StatusCode)
			if test.want.StatusCode == http.StatusOK {
				assert.Equal(tt, test.want.body, body)
			}
		})
	}
}
