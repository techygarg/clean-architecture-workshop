package query

import (
	"context"
	"myapp/operation/dto/response"
	"myapp/operation/mapper"
	"myapp/persistence/provider"
)

type GetTransactionByPaymentCodeQueryHandler struct {
	provider provider.TransactionProvider
}

func NewGetTransactionByPaymentCodeQueryHandler(provider provider.TransactionProvider) GetTransactionByPaymentCodeQueryHandler {
	return GetTransactionByPaymentCodeQueryHandler{
		provider: provider,
	}
}

func (handler GetTransactionByPaymentCodeQueryHandler) Handle(ctx context.Context, query GetTransactionByPaymentCode) (res []response.Transaction, err error) {
	transactions, err := handler.provider.GetTransactionsPaymentMethod(ctx, query.Code)
	if err != nil {
		return res, err
	}
	return mapper.MapToTransactionDtoList(transactions), nil
}
