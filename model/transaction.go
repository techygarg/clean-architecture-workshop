package model

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TransactionModel struct {
	Id                  string `gorm:"primaryKey column:id"`
	RefNum              string `gorm:"column:ref_num"`
	CurrencyCode        string `gorm:"column:currency"`
	Amount              int64  `gorm:"column:amount"`
	Type                string `gorm:"column:type"`
	PaymentMethodCode   string `gorm:"column:payment_method_code"`
	PaymentProviderCode string `gorm:"column:payment_provider_code"`
	UserIdentifier      string `gorm:"column:user_identifier"`
	Status              string `gorm:"column:status"`
}

func (m *TransactionModel) Create(c *gin.Context) error {
	// fetch DB from context which we set up at app start up
	db := c.Value("my-db-instance").(*gorm.DB)
	result := db.Create(m)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
