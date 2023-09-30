package types

type UserIdentifier string
type PaymentCode string
type ProviderCode string
type TransactionStatus string
type TransactionType string
type CurrencyCode string

func (v PaymentCode) ToString() string {
	return string(v)
}

func (v ProviderCode) ToString() string {
	return string(v)
}

func (v TransactionType) ToString() string {
	return string(v)
}

func (v TransactionStatus) ToString() string {
	return string(v)
}

func (v CurrencyCode) ToString() string {
	return string(v)
}
func (v UserIdentifier) ToString() string {
	return string(v)
}
