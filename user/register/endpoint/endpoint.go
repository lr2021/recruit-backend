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
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(model.LoginRequest)
		rsp := response.(model.LoginResponse)
		req.Password = utils.Md5(req.Password)
		if err := auth.CheckReCaptcha(req.Token); err != nil {
			return nil, err
		}
		response, err = userService.InspectUser(model.InspectUserRequest{
			Username: req.Username,
			Tel:      req.Tel,
			Password: req.Password,
		})
		if err != nil {
			return nil, err
		}
		if response == nil {
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
	return nil
}

func Logout(UserService service.IUserService) endpoint.Endpoint {
	return nil
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
	return nil
}

func HealthCheck(userService service.IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		status, _ := userService.HealthCheck(model.ServiceHealthCheckRequest{})
		return model.HealthResponse{Health: status.Health}, nil
	}
}