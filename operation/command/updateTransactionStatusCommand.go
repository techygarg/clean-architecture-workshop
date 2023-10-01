package command

import (
	"context"
	"myapp/operation/externalProvider"
	"myapp/persistence/repository"
)

type UpdateTransactionStatusCommandHandler struct {
	userProvider     externalProvider.UserProvider
	transferProvider externalProvider.TransferProvider
	tranRepo         repository.TransactionRepository
}

func NewUpdateTransactionStatusCommandHandler(repo repository.TransactionRepository) UpdateTransactionStatusCommandHandler {
	return UpdateTransactionStatusCommandHandler{
		tranRepo: repo,
	}
}

func (h UpdateTransactionStatusCommandHandler) Handle(ctx context.Context, command UpdateTransactionStatusCommand) error {
	transaction, err := h.tranRepo.GetTransactionById(ctx, command.TransactionId)
	if err != nil {
		return err
	}
	err = transaction.ChangeStatus(command.Status)
	if err != nil {
		return err
	}

	return h.tranRepo.Update(ctx, *transaction)
}
