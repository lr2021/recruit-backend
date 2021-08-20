package factory

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	httptransport "github.com/go-kit/kit/transport/http"
	"io"
	"net/http"
	"net/url"
	"strings"
)
func encodeRequest(_ context.Context, req *http.Request, r interface{}) error {
	return nil
}

func decodeResponse(_ context.Context, rsp *http.Response) (interface{}, error) {
	return rsp.Body, nil
}

func UserFactory(_ context.Context, method, path string) sd.Factory {
	return func(instance string) (endpoint endpoint.Endpoint, closer io.Closer, err error) {
		if !strings.HasPrefix(instance, "http") {
			instance = "http://" + instance
		}

		target, err := url.Parse(instance)
		if err != nil {
			return nil, nil, err
		}
		target.Path = path
		return httptransport.NewClient(method, target, encodeRequest, decodeResponse).Endpoint(), nil, nil
	}
}
