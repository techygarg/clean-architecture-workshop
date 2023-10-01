package command

import (
	"context"
	"errors"
	"myapp/domain"
	st "myapp/domain/coreTypes"
	"myapp/operation/dto/response"
	"myapp/operation/externalProvider"
	"myapp/persistence/repository"
)

type CreateTransactionCommandHandler struct {
	userProvider     externalProvider.UserProvider
	transferProvider externalProvider.TransferProvider
	tranRepo         repository.TransactionRepository
}

func NewCreateTransactionCommandHandler(user externalProvider.UserProvider,
	transfer externalProvider.TransferProvider, repo repository.TransactionRepository) CreateTransactionCommandHandler {
	return CreateTransactionCommandHandler{
		userProvider:     user,
		transferProvider: transfer,
		tranRepo:         repo,
	}
}

func (h CreateTransactionCommandHandler) Handle(ctx context.Context, command CreateTransactionCommand) (res response.TransactionCreatedResponse, err error) {
	userIdentifier := ctx.Value("user-identifier").(string)

	canDeposit, err := h.userProvider.CanDeposit(userIdentifier)
	if err != nil {
		return res, err
	}
	if !canDeposit {
		return res, errors.New("user can't deposit")
	}

	refNum, err := h.transferProvider.TransferMoneyFromProvider(command)
	if err != nil {
		return res, err
	}

	transaction := domain.Transaction{
		RefNum:              refNum,
		Amount:              command.Amount,
		Type:                st.TransactionType("Credit"),
		PaymentMethodCode:   st.PaymentCode(command.PaymentMethodCode),
		PaymentProviderCode: st.ProviderCode(command.PaymentProviderCode),
		UserIdentifier:      st.UserIdentifier(userIdentifier),
		CurrencyCode:        st.CurrencyCode(command.CurrencyCode),
	}

	err = transaction.ChangeStatus("Initialized")
	if err != nil {
		return res, err
	}

	id, err := h.tranRepo.Create(ctx, transaction)
	res.Id = id
	return res, err
}
