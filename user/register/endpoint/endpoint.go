package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/lr2021/recruit-backend/auth"
	"github.com/lr2021/recruit-backend/general/errors"
	"github.com/lr2021/recruit-backend/user/register/model"
	"github.com/lr2021/recruit-backend/user/register/service"
	"github.com/lr2021/recruit-backend/utils"
)

type Endpoints struct {
	Login endpoint.Endpoint
	Register endpoint.Endpoint
	Logout endpoint.Endpoint
	GetUserProfile endpoint.Endpoint
	GetUserSolved endpoint.Endpoint
	UpdateUserProfile endpoint.Endpoint
	GetUserRank endpoint.Endpoint
	HealthCheck endpoint.Endpoint
}

func Login(userService service.IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var rsp model.LoginResponse

		req := request.(model.LoginRequest)
		req.Password = utils.Md5(req.Password)
		if err := auth.CheckReCaptcha(req.Token); err != nil {
			return nil, err
		}
		response, err := userService.InspectUser(model.InspectUserRequest{
			Username: req.Username,
			Tel:      req.Tel,
			Password: req.Password,
		})
		if err != nil {
			return nil, err
		}
		if response.User == nil {
			return nil, errors.Forbidden("user:login:001", "invalid username or password")
		}
		rsp.Username = req.Username
		rsp.Status = 200
		rsp.Msg = "success"
		rsp.Token = auth.GenerateToken(req.Username)
		return rsp, nil
	}
}

func Register(userService service.IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var rsp model.RegisterResponse

		req := request.(model.RegisterRequest)
		req.Password = utils.Md5(req.Password)

		rspValidation := auth.CheckValidation(req.ValidationCode, req.Tel, req.Username)
		if !rspValidation {
			rsp.Status = 999 // wrong validation code
			rsp.Msg = "wrong validation code"
			return rsp, nil
		}

		response, err := userService.AddUser(model.AddUserRequest{
			User: &model.User{
				Username:            req.Username,
				Password:            req.Password,
				Tel:                 req.Tel,
				StuNumber:           req.StuNumber,
				ProblemSolvedNumber: 0,
				Grade:               req.Grade,
			},
		})

		if err != nil {
			return nil, err
		}
		if response.Status == 1001 {
			rsp.Status = 1001 // tel has registered
			rsp.Msg = "tel has registered"
		}
		if response.Status == 1002 {
			rsp.Status = 1002 // stuNumber has registered
			rsp.Msg = "stuNumber has registered"
		}
		if response.Status == 1003 {
			rsp.Status = 1003 // username has registered
			rsp.Msg = "username has registered"
		}
		if response.Status == 200 {
			rsp.Status = 200 // register success
			rsp.Msg = "success"
		}
		return rsp, nil
	}
}

func Logout(userService service.IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var rsp model.LogoutResponse

		req := request.(model.LogoutRequest)
		err := auth.RemoveToken(req.Username)

		if err != nil {
			return rsp, err
		}
		rsp.Username = req.Username
		rsp.Status = 200
		rsp.Msg = "success"

		return rsp, nil
	}
}

func GetUserProfile(userService service.IUserService) endpoint.Endpoint {
	return nil
}

func UpdateUserProfile(userService service.IUserService) endpoint.Endpoint {
	return nil
}

func GetUserSolved(userService service.IUserService) endpoint.Endpoint {
	return nil
}

func GetUserRank(userService service.IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var rsp model.GetUserRankResponse
		var err error
		rsp.Username = "su29029"
		rsp.Status = 200
		rsp.Msg = "success"
		rsp.Field = "web"
		rsp.Rank = 1
		return rsp, err
	}
}

func HealthCheck(userService service.IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		status, _ := userService.HealthCheck(model.ServiceHealthCheckRequest{})
		return model.HealthResponse{Health: status.Health}, nil
	}
}