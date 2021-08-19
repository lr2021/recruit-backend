package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	httptransport "github.com/go-kit/kit/transport/http"
	"io"
	"net/http"
	"net/url"
	"time"
)

func encodeRequest(_ context.Context, req *http.Request, r interface{}) error {
	return nil
}

func decodeResponse(_ context.Context, rsp *http.Response) (interface{}, error) {
	return rsp.Body, nil
}

func MakeDiscoverEndpoint(ctx context.Context, client consul.Client, logger log.Logger) endpoint.Endpoint {
	serviceName := "lr2021-user"
	tags := []string{"lr2021", "user"}
	passingOnly := true
	duration := 500 * time.Millisecond

	instancer := consul.NewInstancer(client, logger, serviceName, tags, passingOnly)

	factory := func(instance string) (endpoint.Endpoint, io.Closer, error) {
		target, _ := url.Parse("http://" + instance)
		return httptransport.NewClient("", target, encodeRequest, decodeResponse).Endpoint(), nil, nil
	}

	endpointer := sd.NewEndpointer(instancer, factory, logger)

	bl := lb.NewRoundRobin(endpointer)
	retry := lb.Retry(1, duration, bl)

	return retry
}
