package service

import (
	"context"
	"errors"
	"myapp/domain"
	"myapp/persistence/repository"
	"myapp/service/dto/request"
	"myapp/service/dto/response"
	"myapp/service/externalProvider"
)

type TransactionService interface {
	Credit(ctx context.Context, req request.CreateTransactionRequest) (response.TransactionCreatedResponse, error)
}

type transactionService struct {
	userProvider     externalProvider.UserProvider
	transferProvider externalProvider.TransferProvider
	transRepo        repository.TransactionRepository
}

func NewTransactionService(user externalProvider.UserProvider,
	transfer externalProvider.TransferProvider,
	transRepo repository.TransactionRepository) TransactionService {
	return transactionService{
		userProvider:     user,
		transferProvider: transfer,
		transRepo:        transRepo,
	}
}

func (s transactionService) Credit(ctx context.Context, req request.CreateTransactionRequest) (res response.TransactionCreatedResponse, err error) {
	userIdentifier := ctx.Value("user-identifier").(string)

	canDeposit, err := s.userProvider.CanDeposit(userIdentifier)
	if err != nil {
		return res, err
	}
	if !canDeposit {
		return res, errors.New("user can't deposit")
	}

	refNum, err := s.transferProvider.TransferMoneyFromProvider(req)
	if err != nil {
		return res, err
	}

	transaction := domain.Transaction{
		CurrencyCode:        req.CurrencyCode,
		Amount:              req.Amount,
		PaymentMethodCode:   req.PaymentMethodCode,
		PaymentProviderCode: req.PaymentProviderCode,
		Type:                "Credit",
		Status:              "Success",
		UserIdentifier:      userIdentifier,
		RefNum:              refNum,
	}

	id, err := s.transRepo.Create(ctx, transaction)
	res.Id = id
	return res, err
}
