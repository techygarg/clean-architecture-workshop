package externalProvider

import (
	"errors"
	"myapp/service/dto/response"
)

type UserProvider interface {
	CanDeposit(user string) (bool, error)
}

type userProvider struct {
}

func NewUserService() UserProvider {
	return userProvider{}
}

func (p userProvider) CanDeposit(user string) (bool, error) {
	userRes, err := p.getUser(user)
	if err != nil {
		return false, err
	}
	if !userRes.IsActive || !userRes.CanDeposit {
		return false, errors.New("user can initiate debit")
	}
	return true, nil
}

func (p userProvider) getUser(user string) (response.UserResponse, error) {
	return response.UserResponse{}, nil
}
