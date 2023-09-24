package response

type UserResponse struct {
	IsActive    bool
	CanDeposit  bool
	CanWithdraw bool
}
