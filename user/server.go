package main

import (
	"context"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/lr2021/recruit-backend/user/endpoint"
	"github.com/lr2021/recruit-backend/user/transport"
	"net/http"
)

func NewHTTPServer(ctx context.Context, endpoints endpoint.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(authMiddleWare)

	s := r.PathPrefix("/api/user").Subrouter()

	s.Methods("POST").Path("/login").Handler(httptransport.NewServer(
		endpoints.Login,
		transport.DecodeLoginRequest,
		transport.Encode,
	))
	s.Methods("POST").Path("/register").Handler(httptransport.NewServer(
		endpoints.Register,
		transport.DecodeRegisterRequest,
		transport.Encode,
	))
	s.Methods("POST").Path("/logout").Handler(httptransport.NewServer(
		endpoints.Logout,
		transport.DecodeLogoutRequest,
		transport.Encode,
	))
	s.Methods("GET").Path("/userSolved").Handler(httptransport.NewServer(
		endpoints.GetUserSolved,
		transport.DecodeGetUserSolvedRequest,
		transport.Encode,
		))
	s.Methods("GET").Path("/userProfile").Handler(httptransport.NewServer(
		endpoints.GetUserProfile,
		transport.DecodeGetUserProfileRequest,
		transport.Encode,
	))
	s.Methods("PUT").Path("/userProfile").Handler(httptransport.NewServer(
		endpoints.UpdateUserProfile,
		transport.DecodeUpdateUserProfileRequest,
		transport.Encode,
	))
	s.Methods("GEt").Path("/rank").Handler(httptransport.NewServer(
		endpoints.GetUserRank,
		transport.DecodeGetUserRankRequest,
		transport.Encode,
	))

	return r
}

func authMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}