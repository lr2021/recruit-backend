package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/lr2021/recruit-backend/user/model"
	"github.com/lr2021/recruit-backend/user/service"
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
}

func Login(userService service.IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req :=request.(model.LoginRequest)
		req.Password = utils.Md5(req.Password)
		return nil, nil
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