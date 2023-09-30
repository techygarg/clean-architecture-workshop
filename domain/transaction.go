package domain

type Transaction struct {
	Id                  int
	RefNum              string
	CurrencyCode        string
	Amount              int64
	Type                string
	PaymentMethodCode   string
	PaymentProviderCode string
	UserIdentifier      string
	Status              string
}
