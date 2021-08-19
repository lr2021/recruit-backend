package factory

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"io"
	"net/http"
)

func UserFactory(_ context.Context, method, path string) sd.Factory {
	return func(instance string) (endpoint endpoint.Endpoint, closer io.Closer, err error) {
		http.Request{}
	}
}
