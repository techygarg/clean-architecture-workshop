package service

import (
	"context"
	"errors"
	"myapp/domain"
	st "myapp/domain/types"
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
		RefNum:              refNum,
		Amount:              req.Amount,
		Type:                st.TransactionType("Credit"),
		Status:              st.TransactionStatus("Success"),
		PaymentMethodCode:   st.PaymentCode(req.PaymentMethodCode),
		PaymentProviderCode: st.ProviderCode(req.PaymentProviderCode),
		UserIdentifier:      st.UserIdentifier(userIdentifier),
		CurrencyCode:        st.CurrencyCode(req.CurrencyCode),
	}

	id, err := s.transRepo.Create(ctx, transaction)
	res.Id = id
	return res, err
}
