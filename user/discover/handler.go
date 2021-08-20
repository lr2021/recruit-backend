package discover

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/lr2021/recruit-backend/user/discover/transport"
	"net/http"
)

func NewHTTPHandler(ctx context.Context, endpoints map[string]endpoint.Endpoint) http.Handler {
	r := mux.NewRouter()

	s := r.PathPrefix("/api/user").Subrouter()

	s.Methods("POST").Path("/login").Handler(httptransport.NewServer(
		endpoints["Login"],
		transport.DecodeLoginRequest,
		transport.Encode,
	))
	s.Methods("POST").Path("/register").Handler(httptransport.NewServer(
		endpoints["Register"],
		transport.DecodeRegisterRequest,
		transport.Encode,
	))
	s.Methods("POST").Path("/logout").Handler(httptransport.NewServer(
		endpoints["Logout"],
		transport.DecodeLogoutRequest,
		transport.Encode,
	))
	s.Methods("GET").Path("/userSolved").Handler(httptransport.NewServer(
		endpoints["GetUserSolved"],
		transport.DecodeGetUserSolvedRequest,
		transport.Encode,
	))
	s.Methods("GET").Path("/userProfile").Handler(httptransport.NewServer(
		endpoints["GetUserProfile"],
		transport.DecodeGetUserProfileRequest,
		transport.Encode,
	))
	s.Methods("PUT").Path("/userProfile").Handler(httptransport.NewServer(
		endpoints["UpdateUserProfile"],
		transport.DecodeUpdateUserProfileRequest,
		transport.Encode,
	))
	s.Methods("GET").Path("/rank").Handler(httptransport.NewServer(
		endpoints["GetUserRank"],
		transport.DecodeGetUserRankRequest,
		transport.Encode,
	))

	return r
}

