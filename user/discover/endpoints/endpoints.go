package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	httptransport "github.com/go-kit/kit/transport/http"
	factory2 "github.com/lr2021/recruit-backend/user/discover/factory"
	"time"
)

type api struct {
	method string
	path string
}


func MakeDiscoverEndpoint(ctx context.Context, client consul.Client, logger log.Logger) map[string]endpoint.Endpoint {
	serviceName := "lr2021-user"
	tags := []string{"lr2021", "user"}
	passingOnly := true
	duration := 500 * time.Millisecond

	instancer := consul.NewInstancer(client, logger, serviceName, tags, passingOnly)

	endpoints := map[string]endpoint.Endpoint{}
	apis := map[string]api{
		"Login": {"POST", "/api/user/login"},
		"Register": {"POST", "/api/user/register"},
		"Logout": {"POST", "/api/user/logout"},
		"GetUserSolved": {"GET", "/api/user/userSolved"},
		"GetUserProfile": {"GET", "/api/user/userProfile"},
		"UpdateUserProfile": {"PUT", "/api/user/userProfile"},
		"GetUserRank": {"GET", "/api/user/rank"},
	}
	encode := map[string]httptransport.EncodeRequestFunc{}
	decode := map[string]httptransport.DecodeResponseFunc{}

	encode["Login"] = factory2.EncodeLoginRequest
	encode["Register"] = factory2.EncodeRegisterRequest
	encode["Logout"] = factory2.EncodeLogoutRequest
	encode["GetUserSolved"] = factory2.EncodeGetUserSolvedRequest
	encode["GetUserProfile"] = factory2.EncodeGetUserProfileRequest
	encode["UpdateUserProfile"] = factory2.EncodeUpdateUserProfileRequest
	encode["GetUserRank"] = factory2.EncodeGetUserRankRequest

	decode["Login"] = factory2.DecodeLoginResponse
	decode["Register"] = factory2.DecodeRegisterResponse
	decode["Logout"] = factory2.DecodeLogoutResponse
	decode["GetUserSolved"] = factory2.DecodeGetUserSolvedResponse
	decode["GetUserProfile"] = factory2.DecodeGetUserProfileResponse
	decode["UpdateUserProfile"] = factory2.DecodeUpdateUserProfileResponse
	decode["GetUserRank"] = factory2.DecodeGetUserRankResponse

	for c, u := range apis {
		factory := factory2.UserFactory(ctx, encode[c], decode[c], u.path, u.method)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		bl := lb.NewRoundRobin(endpointer)
		endp := lb.Retry(3, duration, bl)
		endpoints[c] = endp
	}

	return endpoints
}
