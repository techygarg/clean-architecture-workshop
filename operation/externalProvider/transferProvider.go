package externalProvider

type TransferProvider interface {
	TransferMoneyFromProvider(req any) (string, error)
}

type transferProvider struct {
}

func NewTransferService() TransferProvider {
	return transferProvider{}
}

func (p transferProvider) TransferMoneyFromProvider(req any) (string, error) {
	// used any just to reduce code changes, ideally it should a proper type
	return "", nil
}
