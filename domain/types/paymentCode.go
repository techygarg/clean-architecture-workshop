package types

const (
	PcUPI        PaymentCode = "UPI"
	PcEWallet    PaymentCode = "EWallet"
	PcCreditCard PaymentCode = "CreditCard"
)

func (v PaymentCode) ToString() string {
	return string(v)
}

func (v PaymentCode) IsUpi() bool {
	return v == PcUPI
}
