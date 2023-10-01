package provider

import (
	"context"
	st "myapp/domain/coreTypes"
	"myapp/persistence/dao"

	"gorm.io/gorm"
)

type TransactionProvider interface {
	GetTransactionById(ctx context.Context, id int) (res dao.TransactionDao, err error)
	GetTransactionByRef(ctx context.Context, refNum string) (res dao.TransactionDao, err error)
	GetTransactionsPaymentMethod(ctx context.Context, code st.PaymentCode) (res []dao.TransactionDao, err error)
	GetTransactionsPaymentProvider(ctx context.Context, code st.ProviderCode) (res []dao.TransactionDao, err error)
}

type transactionProvider struct {
	db *gorm.DB
}

func NewTransactionProvider(db *gorm.DB) TransactionProvider {
	return transactionProvider{db: db}
}

func (r transactionProvider) GetTransactionById(ctx context.Context, id int) (res dao.TransactionDao, err error) {
	return res, err
}

func (r transactionProvider) GetTransactionByRef(ctx context.Context, refNum string) (res dao.TransactionDao, err error) {
	return res, err
}

func (r transactionProvider) GetTransactionsPaymentMethod(ctx context.Context, code st.PaymentCode) (res []dao.TransactionDao, err error) {
	return res, err
}

func (r transactionProvider) GetTransactionsPaymentProvider(ctx context.Context, code st.ProviderCode) (res []dao.TransactionDao, err error) {
	return res, err
}
