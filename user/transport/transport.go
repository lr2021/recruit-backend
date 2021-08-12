package transport

import (
	"context"
	"encoding/json"
	"github.com/lr2021/recruit-backend/general/errors"
	"github.com/lr2021/recruit-backend/user/model"
	"net/http"
)

func DecodeLoginRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.BadRequest("para:001", "missing or invalid parameters")
	}
	return req, nil
}

func DecodeRegisterRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.BadRequest("para:001", "missing or invalid parameters")
	}
	return req, nil
}

func DecodeLogoutRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.LogoutRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.BadRequest("para:001", "missing or invalid parameters")
	}
	return req, nil
}

func DecodeGetUserSolvedRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.GetUserSolvedRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.BadRequest("para:001", "missing or invalid parameters")
	}
	return req, nil
}

func DecodeGetUserProfileRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.GetUserProfileRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.BadRequest("para:001", "missing or invalid parameters")
	}
	return req, nil
}

func DecodeUpdateUserProfileRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.UpdateUserProfileRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.BadRequest("para:001", "missing or invalid parameters")
	}
	return req, nil
}


func DecodeGetUserRankRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.UpdateUserProfileRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.BadRequest("para:001", "missing or invalid parameters")
	}
	return req, nil
}

func Encode(ctx context.Context, w http.ResponseWriter, rsp interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(rsp)
}
