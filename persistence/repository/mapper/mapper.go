package mapper

import (
	"myapp/domain"
	"myapp/persistence/dao"
)

func ToTransactionDao(tran domain.Transaction) dao.TransactionDao {
	return dao.TransactionDao{
		Id:                  tran.Id,
		RefNum:              tran.RefNum,
		CurrencyCode:        tran.CurrencyCode.ToString(),
		Amount:              tran.Amount,
		Type:                tran.Type.ToString(),
		PaymentMethodCode:   tran.PaymentMethodCode.ToString(),
		PaymentProviderCode: tran.PaymentProviderCode.ToString(),
		UserIdentifier:      tran.UserIdentifier.ToString(),
		Status:              tran.Status.ToString(),
	}
}
