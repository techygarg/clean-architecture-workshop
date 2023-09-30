package repository

import (
	"context"
	"myapp/domain"
	"myapp/persistence/repository/mapper"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(ctx context.Context, tran domain.Transaction) (int, error)
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
