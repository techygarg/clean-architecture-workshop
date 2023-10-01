package service

import (
	"context"
	"errors"
	"myapp/domain"
	st "myapp/domain/coreTypes"
	"myapp/persistence/repository"
	"myapp/service/dto/request"
	"myapp/service/dto/response"
	"myapp/service/externalProvider"
	"myapp/service/mapper"
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

func (s transactionService) GetTransactionById(ctx context.Context, id int) (res response.Transaction, err error) {
	transaction, err := s.transRepo.GetTransactionById(ctx, id)
	if err != nil {
		return res, err
	}
	return mapper.MapTransactionDomainToDto(transaction), nil
}

func (s transactionService) GetTransactionByRef(ctx context.Context, refNum string) (res response.Transaction, err error) {
	transaction, err := s.transRepo.GetTransactionByRef(ctx, refNum)
	if err != nil {
		return res, err
	}
	return mapper.MapTransactionDomainToDto(transaction), nil
}

func (s transactionService) GetTransactionsPaymentMethod(ctx context.Context, code st.PaymentCode) (res response.TransactionList, err error) {
	transactions, err := s.transRepo.GetTransactionsPaymentMethod(ctx, code)
	if err != nil {
		return res, err
	}
	return mapper.MapTransactionDomainsToDto(transactions), nil
}

func (s transactionService) GetTransactionsPaymentProvider(ctx context.Context, code st.ProviderCode) (res response.TransactionList, err error) {
	transactions, err := s.transRepo.GetTransactionsPaymentProvider(ctx, code)
	if err != nil {
		return res, err
	}
	return mapper.MapTransactionDomainsToDto(transactions), nil
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
		PaymentMethodCode:   st.PaymentCode(req.PaymentMethodCode),
		PaymentProviderCode: st.ProviderCode(req.PaymentProviderCode),
		UserIdentifier:      st.UserIdentifier(userIdentifier),
		CurrencyCode:        st.CurrencyCode(req.CurrencyCode),
	}

	err = transaction.ChangeStatus("Initialized")
	if err != nil {
		return res, err
	}

	id, err := s.transRepo.Create(ctx, transaction)
	res.Id = id
	return res, err
}

func (s transactionService) Debit(context.Context) error {
	return nil
}

func (s transactionService) Update(context.Context) error {
	return nil
}

func (s transactionService) UpdateStatus(context.Context) error {
	return nil
}
