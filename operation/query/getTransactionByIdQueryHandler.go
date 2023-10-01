package query

import (
	"context"
	"myapp/operation/dto/response"
	"myapp/operation/mapper"
	"myapp/persistence/provider"
)

type GetTransactionByIdQueryHandler struct {
	provider provider.TransactionProvider
}

func NewGetTransactionQueryHandler(provider provider.TransactionProvider) GetTransactionByIdQueryHandler {
	return GetTransactionByIdQueryHandler{
		provider: provider,
	}
}

func (handler GetTransactionByIdQueryHandler) Handle(ctx context.Context, query GetTransactionByIdQuery) (res response.Transaction, err error) {
	transaction, err := handler.provider.GetTransactionById(ctx, query.Id)
	if err != nil {
		return res, err
	}
	return mapper.MapToTransactionDto(transaction), nil
}
