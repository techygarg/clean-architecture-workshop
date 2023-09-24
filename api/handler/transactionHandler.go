package handler

import (
	"myapp/service"
	request2 "myapp/service/dto/request"
	"myapp/service/externalProvider"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService service.TransactionService
}

func NewTransactionHandler(transactionService service.TransactionService, transfer externalProvider.TransferProvider) transactionHandler {
	return transactionHandler{
		transactionService: transactionService,
	}
}
func (h transactionHandler) Credit(c *gin.Context) {

	req := request2.CreateTransactionRequest{}
	if err := c.BindJSON(&req); err != nil {
		return
	}

	err := h.transactionService.Credit(c, req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, "transaction created")
}
