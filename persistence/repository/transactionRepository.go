package repository

import (
	"context"
	"myapp/domain"
	st "myapp/domain/coreTypes"
	"myapp/persistence/repository/mapper"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(ctx context.Context, tran domain.Transaction) (int, error)
	Update(ctx context.Context, tran domain.Transaction) error
	GetTransactionById(ctx context.Context, id int) (res *domain.Transaction, err error)
	GetTransactionByRef(ctx context.Context, refNum string) (res *domain.Transaction, err error)
	GetTransactionsPaymentMethod(ctx context.Context, code st.PaymentCode) (res []*domain.Transaction, err error)
	GetTransactionsPaymentProvider(ctx context.Context, code st.ProviderCode) (res []*domain.Transaction, err error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return transactionRepository{db: db}
}

func (r transactionRepository) Create(ctx context.Context, tran domain.Transaction) (id int, err error) {
	tranDao := mapper.ToTransactionDao(tran)
	result := r.db.Create(tranDao)
	if result.Error != nil {
		return id, result.Error
	}
	return tran.Id, nil
}

func (r transactionRepository) Update(ctx context.Context, tran domain.Transaction) error {
	tranDao := mapper.ToTransactionDao(tran)
	result := r.db.Save(tranDao)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r transactionRepository) GetTransactionById(ctx context.Context, id int) (res *domain.Transaction, err error) {
	return res, err
}

func (r transactionRepository) GetTransactionByRef(ctx context.Context, refNum string) (res *domain.Transaction, err error) {
	return res, err
}

func (r transactionRepository) GetTransactionsPaymentMethod(ctx context.Context, code st.PaymentCode) (res []*domain.Transaction, err error) {
	return res, err
}

func (r transactionRepository) GetTransactionsPaymentProvider(ctx context.Context, code st.ProviderCode) (res []*domain.Transaction, err error) {
	return res, err
}
