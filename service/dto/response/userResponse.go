package response

type UserResponse struct {
	IsActive    bool
	CanDeposit  bool
	CanWithdraw bool
}

type TransactionCreatedResponse struct {
	Id int `json:"id"`
}

type Transaction struct {
}

type TransactionList struct {
	Transactions []Transaction `json:"transactions"`
}
