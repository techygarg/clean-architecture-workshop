package externalProvider

import (
	"myapp/service/dto/response"
)

type UserProvider interface {
	GetUser(user string) (response.UserResponse, error)
}

type userProvider struct {
}

func NewUserService() UserProvider {
	return userProvider{}
}

func (p userProvider) GetUser(user string) (response.UserResponse, error) {
	return response.UserResponse{}, nil
}
