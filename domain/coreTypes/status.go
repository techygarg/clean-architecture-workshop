package coreTypes

const (
	TrStatusInitialized TransactionStatus = "Initialized"
	TrStatusInProgress  TransactionStatus = "InProgress"
	TrStatusCredited    TransactionStatus = "Credited"
	TrStatusDebited     TransactionStatus = "Debited"
	TrStatusFailed      TransactionStatus = "Failed"
)

var StateTransitionMap = map[TransactionStatus][]TransactionStatus{
	TrStatusInitialized: {TrStatusInProgress, TrStatusFailed},
	TrStatusInProgress:  {TrStatusCredited, TrStatusDebited, TrStatusFailed},
}

func (v TransactionStatus) ToString() string {
	return string(v)
}

func (v TransactionStatus) IsInitialized() bool {
	return v == TrStatusInitialized
}

func (v TransactionStatus) IsCredited() bool {
	return v == TrStatusCredited
}

func (v TransactionStatus) CanTransitionTo(status TransactionStatus) bool {
	// all custom logic
	return false
}
