package service

import (
	"github.com/wen-qu/kit-xuesou-backend/user/model"
)

type IUserService interface {
	InspectUser(req model.InspectRequest) (model.InspectResponse, error)
	UpdateUser(req model.UpdateRequest) (model.UpdateResponse, error)
	AddUser(req model.AddRequest) (model.AddResponse, error)
	DeleteUser(req model.DeleteRequest) (model.DeleteResponse, error)
}

type userService struct {}

func NewService() IUserService {
	return userService{}
}