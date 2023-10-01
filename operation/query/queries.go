package query

import "myapp/domain/coreTypes"

type GetTransactionByIdQuery struct {
	Id int `uri:"id" binding:"required"`
}

type GetTransactionByPaymentCode struct {
	Code coreTypes.PaymentCode `uri:"code" binding:"required"`
}
