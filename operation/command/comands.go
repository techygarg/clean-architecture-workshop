package command

import "myapp/domain/coreTypes"

type CreateTransactionCommand struct {
	CurrencyCode        string `json:"currency"`
	Amount              int64  `json:"amount"`
	Type                string `json:"type"`
	PaymentMethodCode   string `json:"paymentMethodCode"`
	PaymentProviderCode string `json:"paymentProviderCode"`
}

type UpdateTransactionStatusCommand struct {
	TransactionId int                         `json:"transactionId"`
	Status        coreTypes.TransactionStatus `json:"status"`
}
