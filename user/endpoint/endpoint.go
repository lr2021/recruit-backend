package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/lr2021/recruit-backend/user/service"
)

type Endpoints struct {
	Login endpoint.Endpoint
	Register endpoint.Endpoint
	Logout endpoint.Endpoint
	GetUserInformation endpoint.Endpoint
	GetUserSolves endpoint.Endpoint
	UpdateUserProfile endpoint.Endpoint
	GetAllUserProfile endpoint.Endpoint
}

func Login(userService service.IUserService) endpoint.Endpoint {
	return nil
}

func Register(userService service.IUserService) endpoint.Endpoint {
	return nil
}

func Logout(UserService service.IUserService) endpoint.Endpoint {
	return nil
}

func GetUserInformation(userService service.IUserService) endpoint.Endpoint {
	return nil
}

func GetUserSolves(userService service.IUserService) endpoint.Endpoint {
	return nil
}

func UpdateUserProfile(userService service.IUserService) endpoint.Endpoint {
	return nil
}

func GetAllUserProfile(userService service.IUserService) endpoint.Endpoint {
	return nil
}