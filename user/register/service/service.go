package service

import (
	"github.com/lr2021/recruit-backend/user/register/model"
)

type IUserService interface {
	InspectUser(req model.InspectUserRequest) (model.InspectUserResponse, error)
	UpdateUser(req model.UpdateUserRequest) (model.UpdateUserResponse, error)
	AddUser(req model.AddUserRequest) (model.AddUserResponse, error)
	DeleteUser(req model.DeleteUserRequest) (model.DeleteUserResponse, error)
	HealthCheck(req model.ServiceHealthCheckRequest) (model.ServiceHealthCheckResponse, error)
}

type userService struct {}

func NewService() IUserService {
	return userService{}
}

func (u userService) InspectUser(req model.InspectUserRequest) (model.InspectUserResponse, error) {
	return model.InspectUserResponse{}, nil
}

func (u userService) UpdateUser(req model.UpdateUserRequest) (model.UpdateUserResponse, error) {
	return model.UpdateUserResponse{}, nil
}

func (u userService) AddUser(req model.AddUserRequest) (model.AddUserResponse, error) {
	return model.AddUserResponse{}, nil
}

func (u userService) DeleteUser(req model.DeleteUserRequest) (model.DeleteUserResponse, error) {
	return model.DeleteUserResponse{}, nil
}

func (u userService) HealthCheck(req model.ServiceHealthCheckRequest) (model.ServiceHealthCheckResponse, error) {
	return model.ServiceHealthCheckResponse{Health: true}, nil
}
