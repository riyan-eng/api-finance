package repositories

import (
	"fmt"

	"github.com/riyan-eng/api-finance/src/finance/repositories/models"
	"github.com/riyan-eng/api-finance/src/finance/services/entities"
	"gorm.io/gorm"
)

type CashRepositoryInterface interface {
	CashReceipt(cashEntity entities.CashReceipt) error
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

func (cash *cashRepository) CashReceipt(cashEntity entities.CashReceipt) error {

	// debet
	var cashModelDebet models.CashModel = models.CashModel{
		ID:          "lalala",
		UserID:      cashEntity.UserID,
		Code:        cashEntity.Journal.Debet.Code,
		Name:        cashEntity.Journal.Debet.Name,
		Amount:      cashEntity.Journal.Debet.Amount,
		Description: cashEntity.Journal.Debet.Description,
	}
	queryDebet := fmt.Sprintf(`
		INSERT INTO finance.general_ledgers (id, coa, description, debet, user_id) VALUES (%v, %v, %v, %v, %v)
	`, cashModelDebet.ID, cashModelDebet.Code, cashModelDebet.Description, cashModelDebet.Amount, cashModelDebet.UserID)
	cash.DB.Exec(queryDebet)

	// credit
	var cashModelCredit models.CashModel = models.CashModel{
		ID:          "lalala",
		UserID:      cashEntity.UserID,
		Code:        cashEntity.Journal.Debet.Code,
		Name:        cashEntity.Journal.Debet.Name,
		Amount:      cashEntity.Journal.Debet.Amount,
		Description: cashEntity.Journal.Debet.Description,
	}
	queryCredit := fmt.Sprintf(`
		INSERT INTO finance.general_ledgers (id, coa, description, credit, user_id) VALUES (%v, %v, %v, %v, %v)
	`, cashModelCredit.ID, cashModelCredit.Code, cashModelCredit.Description, cashModelCredit.Amount, cashModelCredit.UserID)
	cash.DB.Exec(queryCredit)

	return nil
}
