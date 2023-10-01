package mapper

import (
	"myapp/domain"
	"myapp/service/dto/response"
)

func MapTransactionDomainToDto(tran *domain.Transaction) (res response.Transaction) {
	// map all required fields
	return res
}

func MapTransactionDomainsToDto(tran []*domain.Transaction) (res response.TransactionList) {
	// map all required fields
	return res
}
