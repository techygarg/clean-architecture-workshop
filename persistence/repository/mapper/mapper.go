package mapper

import (
	"myapp/domain"
	"myapp/persistence/dao"
)

func ToTransactionDao(tran domain.Transaction) dao.TransactionDao {
	return dao.TransactionDao{
		Id:                  tran.Id,
		RefNum:              tran.RefNum,
		CurrencyCode:        tran.CurrencyCode,
		Amount:              tran.Amount,
		Type:                tran.Type,
		PaymentMethodCode:   tran.PaymentMethodCode,
		PaymentProviderCode: tran.PaymentProviderCode,
		UserIdentifier:      tran.UserIdentifier,
		Status:              tran.Status,
	}
}
