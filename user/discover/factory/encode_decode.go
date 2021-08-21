package factory

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/lr2021/recruit-backend/user/register/model"
	"io/ioutil"
	"net/http"
)

func EncodeLoginRequest(_ context.Context, req *http.Request, r interface{}) error {
	request := r.(model.LoginRequest)
	bodyBytes, _ := json.Marshal(request)

	req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	return nil
}

func EncodeRegisterRequest(_ context.Context, req *http.Request, r interface{}) error {
	request := r.(model.RegisterRequest)
	bodyBytes, _ := json.Marshal(request)

	req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	return nil
}

func EncodeLogoutRequest(_ context.Context, req *http.Request, r interface{}) error {
	request := r.(model.LogoutRequest)
	bodyBytes, _ := json.Marshal(request)

	req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	return nil
}

func EncodeGetUserSolvedRequest(_ context.Context, req *http.Request, r interface{}) error {
	request := r.(model.GetUserSolvedRequest)
	bodyBytes, _ := json.Marshal(request)

	req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	return nil
}

func EncodeGetUserProfileRequest(_ context.Context, req *http.Request, r interface{}) error {
	request := r.(model.GetUserProfileRequest)
	bodyBytes, _ := json.Marshal(request)

	req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	return nil
}

func EncodeUpdateUserProfileRequest(_ context.Context, req *http.Request, r interface{}) error {
	request := r.(model.UpdateUserProfileRequest)
	bodyBytes, _ := json.Marshal(request)

	req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	return nil
}

func EncodeGetUserRankRequest(_ context.Context, req *http.Request, r interface{}) error {
	request := r.(model.GetUserRankRequest)
	bodyBytes, _ := json.Marshal(request)

	req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	return nil
}

func DecodeLoginResponse(_ context.Context, rsp *http.Response) (interface{}, error) {
	var response model.LoginResponse
	if err := json.NewDecoder(rsp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response, nil
}

func DecodeRegisterResponse(_ context.Context, rsp *http.Response) (interface{}, error) {
	var response model.RegisterResponse
	if err := json.NewDecoder(rsp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response, nil
}

func DecodeLogoutResponse(_ context.Context, rsp *http.Response) (interface{}, error) {
	var response model.LoginResponse
	if err := json.NewDecoder(rsp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response, nil
}

func DecodeGetUserSolvedResponse(_ context.Context, rsp *http.Response) (interface{}, error) {
	var response model.GetUserSolvedResponse
	if err := json.NewDecoder(rsp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response, nil
}

func DecodeGetUserProfileResponse(_ context.Context, rsp *http.Response) (interface{}, error) {
	var response model.GetUserProfileResponse
	if err := json.NewDecoder(rsp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response, nil
}

func DecodeUpdateUserProfileResponse(_ context.Context, rsp *http.Response) (interface{}, error) {
	var response model.UpdateUserProfileResponse
	if err := json.NewDecoder(rsp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response, nil
}

func DecodeGetUserRankResponse(_ context.Context, rsp *http.Response) (interface{}, error) {
	var response model.GetUserRankResponse
	if err := json.NewDecoder(rsp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response, nil
}