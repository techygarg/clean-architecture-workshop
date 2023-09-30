package dao

type TransactionModel struct {
	Id                  int    `gorm:"primaryKey column:id"`
	RefNum              string `gorm:"column:ref_num"`
	CurrencyCode        string `gorm:"column:currency"`
	Amount              int64  `gorm:"column:amount"`
	Type                string `gorm:"column:type"`
	PaymentMethodCode   string `gorm:"column:payment_method_code"`
	PaymentProviderCode string `gorm:"column:payment_provider_code"`
	UserIdentifier      string `gorm:"column:user_identifier"`
	Status              string `gorm:"column:status"`
}
