package repository

import (
	"context"
	"myapp/persistence/dao"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(ctx context.Context, tran dao.TransactionModel) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return transactionRepository{db: db}
}

func (r transactionRepository) Create(ctx context.Context, tran dao.TransactionModel) error {
	result := r.db.Create(tran)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
