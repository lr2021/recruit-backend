package endpoint

import "github.com/go-kit/kit/endpoint"

type Endpoints struct {
	Login endpoint.Endpoint
	Register endpoint.Endpoint
	UpdateProfile endpoint.Endpoint
	ReadProfile endpoint.Endpoint
}