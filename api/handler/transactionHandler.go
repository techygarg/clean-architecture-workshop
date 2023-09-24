package handler

import (
	"errors"
	"myapp/model"
	request2 "myapp/model/request"
	"myapp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	user     service.UserService
	transfer service.TransferService
}

func NewTransactionHandler(user service.UserService, transfer service.TransferService) transactionHandler {
	return transactionHandler{
		user:     user,
		transfer: transfer,
	}
}
func (h transactionHandler) Credit(c *gin.Context) {

	req := request2.CreateTransactionRequest{}
	if err := c.BindJSON(&req); err != nil {
		return
	}

	userIdentifier := c.GetString("user-identifier")

	user, _ := h.user.GetUser(userIdentifier)
	if !user.IsActive || !user.CanDeposit {
		_ = c.Error(errors.New("user can initiate debit"))
		return
	}

	refNum, err := h.transfer.TransferMoneyFromProvider(req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	transaction := model.TransactionModel{
		CurrencyCode:        req.CurrencyCode,
		Amount:              req.Amount,
		PaymentMethodCode:   req.PaymentMethodCode,
		PaymentProviderCode: req.PaymentProviderCode,
		Type:                "Credit",
		Status:              "Success",
		UserIdentifier:      userIdentifier,
		RefNum:              refNum,
	}

	err = transaction.Create(c)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, "transaction created")
}
