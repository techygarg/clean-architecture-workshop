package externalProvider

import (
	request2 "myapp/service/dto/request"
)

type TransferProvider interface {
	TransferMoneyFromProvider(req request2.CreateTransactionRequest) (string, error)
}

type transferProvider struct {
}

func NewTransferService() TransferProvider {
	return transferProvider{}
}

func (p transferProvider) TransferMoneyFromProvider(req request2.CreateTransactionRequest) (string, error) {
	return "", nil
}
