package service

import (
	"context"
	"errors"
	"myapp/persistence/dao"
	"myapp/persistence/repository"
	"myapp/service/dto/request"
	"myapp/service/externalProvider"
)

type TransactionService interface {
	Credit(ctx context.Context, req request.CreateTransactionRequest) error
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

func (s transactionService) Credit(ctx context.Context, req request.CreateTransactionRequest) error {
	userIdentifier := ctx.Value("user-identifier").(string)

	canDeposit, err := s.userProvider.CanDeposit(userIdentifier)
	if err != nil {
		return err
	}
	if !canDeposit {
		return errors.New("user can't deposit")
	}

	refNum, err := s.transferProvider.TransferMoneyFromProvider(req)
	if err != nil {
		return err
	}

	transaction := dao.TransactionModel{
		CurrencyCode:        req.CurrencyCode,
		Amount:              req.Amount,
		PaymentMethodCode:   req.PaymentMethodCode,
		PaymentProviderCode: req.PaymentProviderCode,
		Type:                "Credit",
		Status:              "Success",
		UserIdentifier:      userIdentifier,
		RefNum:              refNum,
	}

	return s.transRepo.Create(ctx, transaction)
}
