package request

type CreateTransactionRequest struct {
	CurrencyCode        string `json:"column:currency"`
	Amount              int64  `json:"column:amount"`
	Type                string `json:"column:type"`
	PaymentMethodCode   string `json:"column:paymentMethodCode"`
	PaymentProviderCode string `json:"column:paymentProviderCode"`
}
