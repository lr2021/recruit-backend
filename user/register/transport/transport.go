package transport

import (
	"context"
	"encoding/json"
	"github.com/lr2021/recruit-backend/general/errors"
	"github.com/lr2021/recruit-backend/user/register/model"
	"net/http"
	"regexp"
)

func CheckLoginRequest(req model.LoginRequest) bool {
	if len(req.Username) == 0 && len(req.Tel) == 0 {
		return false
	}
	if len(req.Username) > 0 && !regexp.MustCompile(`^[0-9a-zA-Z]{6,20}$`).MatchString(req.Username) {
		return false
	}
	if !regexp.MustCompile(`^[0-9a-zA-Z]{6,20}$`).MatchString(req.Password) {
		return false
	}
	if len(req.Tel) > 0 && !regexp.MustCompile(`^((0\d{2,3}-\d{7,8})|(1[358479]\d{9}))$`).MatchString(req.Tel) {
		return false
	}
	return true
}

func CheckRegisterRequest(req model.RegisterRequest) bool {
	if req.Password != req.RepeatedPassword {
		return false
	}
	if !regexp.MustCompile(`^[0-9a-zA-Z]{6,20}$`).MatchString(req.Username) {
		return false
	}
	if !regexp.MustCompile(`^[0-9a-zA-Z]{6,20}$`).MatchString(req.Password) {
		return false
	}
	if !regexp.MustCompile(`^((0\d{2,3}-\d{7,8})|(1[358479]\d{9}))$`).MatchString(req.Tel) {
		return false
	}

	return true
}

func CheckLogoutRequest(req model.LogoutRequest) bool {
	if !regexp.MustCompile(`^[0-9a-zA-Z]{6,20}$`).MatchString(req.Username) {
		return false
	}
	return true
}

func CheckGetUserSolvedRequest(req model.GetUserSolvedRequest) bool {
	if !regexp.MustCompile(`^[0-9a-zA-Z]{6,20}$`).MatchString(req.Username) {
		return false
	}
	return true
}

func CheckGetUserProfileRequest(req model.GetUserProfileRequest) bool {
	if len(req.Tel) == 0 && len(req.Username) == 0 {
		return false
	}
	if len(req.Username) > 0 && !regexp.MustCompile(`^[0-9a-zA-Z]{6,20}$`).MatchString(req.Username) {
		return false
	}
	if len(req.Tel) > 0 && !regexp.MustCompile(`^((0\d{2,3}-\d{7,8})|(1[358479]\d{9}))$`).MatchString(req.Tel) {
		return false
	}
	if req.Type < 1 || req.Type > 5 {
		return false
	}
	return true
}

func CheckUpdateUserProfileRequest(req model.UpdateUserProfileRequest) bool {
	return true
}

func CheckGetUserRankRequest(req model.GetUserRankRequest) bool {
	return true
}

func DecodeLoginRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || !CheckLoginRequest(req) {
		return nil, errors.BadRequest("para:001", "missing or invalid parameters")
	}
	return req, nil
}

func DecodeRegisterRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || !CheckRegisterRequest(req) {
		return nil, errors.BadRequest("para:001", "missing or invalid parameters")
	}
	return req, nil
}

func DecodeLogoutRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.LogoutRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || !CheckLogoutRequest(req) {
		return nil, errors.BadRequest("para:001", "missing or invalid parameters")
	}
	return req, nil
}

func DecodeGetUserSolvedRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.GetUserSolvedRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || !CheckGetUserSolvedRequest(req) {
		return nil, errors.BadRequest("para:001", "missing or invalid parameters")
	}
	return req, nil
}

func DecodeGetUserProfileRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.GetUserProfileRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || !CheckGetUserProfileRequest(req) {
		return nil, errors.BadRequest("para:001", "missing or invalid parameters")
	}
	return req, nil
}

func DecodeUpdateUserProfileRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.UpdateUserProfileRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || !CheckUpdateUserProfileRequest(req) {
		return nil, errors.BadRequest("para:001", "missing or invalid parameters")
	}
	return req, nil
}

func DecodeGetUserRankRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.GetUserRankRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || !CheckGetUserRankRequest(req) {
		return nil, errors.BadRequest("para:001", "missing or invalid parameters")
	}
	return req, nil
}

func DecodeHealthCheckRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.HealthRequest
	//err := json.NewDecoder(r.Body).Decode(&req)
	//if err != nil {
	//	return nil, errors.BadRequest("para:001", "missing or invalid parameters")
	//}
	return req, nil
}

func Encode(ctx context.Context, w http.ResponseWriter, rsp interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(rsp)
}
