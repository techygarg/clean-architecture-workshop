package domain

import (
	. "myapp/domain/types"
)

type Transaction struct {
	Id                  int
	RefNum              string
	Amount              int64
	CurrencyCode        CurrencyCode
	Type                TransactionType
	PaymentMethodCode   PaymentCode
	PaymentProviderCode ProviderCode
	UserIdentifier      UserIdentifier
	Status              TransactionStatus
}
