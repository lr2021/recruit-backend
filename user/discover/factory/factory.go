package factory

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	httptransport "github.com/go-kit/kit/transport/http"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func decodeResponse(_ context.Context, rsp *http.Response) (interface{}, error) {
	fmt.Println(rsp.Status)
	response, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err.Error(), nil
	}
	return string(response[:len(response)-1]), nil
}

func UserFactory(_ context.Context, enc httptransport.EncodeRequestFunc, dec httptransport.DecodeResponseFunc, path, method string) sd.Factory {
	return func(instance string) (endpoint endpoint.Endpoint, closer io.Closer, err error) {
		if !strings.HasPrefix(instance, "http") {
			instance = "http://" + instance
		}

		target, err := url.Parse(instance)
		if err != nil {
			return nil, nil, err
		}
		target.Path = path

		return httptransport.NewClient(method, target, enc, dec).Endpoint(), nil, nil
	}
}

















