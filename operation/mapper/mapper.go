package mapper

import (
	"myapp/operation/dto/response"
	"myapp/persistence/dao"
)

func MapToTransactionDto(tran dao.TransactionDao) response.Transaction {
	return response.Transaction{}
}

func MapToTransactionDtoList(tran []dao.TransactionDao) []response.Transaction {
	return []response.Transaction{}
}
