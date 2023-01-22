package repositories

import (
	"github.com/riyan-eng/api-finance/src/finance/services/entities"
	"gorm.io/gorm"
)

type CashRepositoryInterface interface {
	CashReceipt(cashEntity entities.CashReceipt)
	// CashPayment()
}

// var db = database.DB

// func Save(product models.CashModel) models.CashModel {
// 	db.Exec(`
// 	INSERT
// 	`)
// 	return product
// }

type cashRepository struct {
	DB *gorm.DB
}

func NewCashRepository(DB *gorm.DB) CashRepositoryInterface {
	return &cashRepository{
		DB: DB,
	}
}

func (cash *cashRepository) CashReceipt(cashEntity entities.CashReceipt) {
	// cash.DB.Raw("").Scan()
}
