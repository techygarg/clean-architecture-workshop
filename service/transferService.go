package service

import (
	request2 "myapp/model/request"
)

type TransferService interface {
	TransferMoneyFromProvider(req request2.CreateTransactionRequest) (string, error)
}

type transferService struct {
}

func NewTransferService() TransferService {
	return transferService{}
}

func (s transferService) TransferMoneyFromProvider(req request2.CreateTransactionRequest) (string, error) {
	return "", nil
}
