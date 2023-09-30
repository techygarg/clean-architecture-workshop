package repository

import (
	"context"
	"myapp/persistence/dao"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(ctx context.Context, tran dao.TransactionModel) (int, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return transactionRepository{db: db}
}

func (r transactionRepository) Create(ctx context.Context, tran dao.TransactionModel) (id int, err error) {
	result := r.db.Create(tran)
	if result.Error != nil {
		return id, result.Error
	}
	return tran.Id, nil
}
