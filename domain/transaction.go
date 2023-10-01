package domain

import (
	"errors"
	. "myapp/domain/coreTypes"
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

func (d *Transaction) ChangeStatus(newStatus TransactionStatus) error {
	if d.Status == "" && newStatus.IsInitialized() {
		d.Status = newStatus
		return nil
	}
	for _, status := range StateTransitionMap[d.Status] {
		if status == newStatus {
			d.Status = newStatus
			return nil
		}
	}
	return errors.New("invalid state transition")
}
