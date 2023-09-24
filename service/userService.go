package service

import (
	"myapp/model/response"
)

type UserService interface {
	GetUser(user string) (response.UserResponse, error)
}

type userService struct {
}

func NewUserService() UserService {
	return userService{}
}

func (s userService) GetUser(user string) (response.UserResponse, error) {
	return response.UserResponse{}, nil
}
