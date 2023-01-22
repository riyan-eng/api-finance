package repositories

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/riyan-eng/api-finance/src/finance/repositories/models"
	"github.com/riyan-eng/api-finance/src/finance/services/entities"
	"gorm.io/gorm"
)

type CashRepository interface {
	CashReceipt(cashEntity entities.CashReceipt) error
	// CashPayment()
}

type cashRepository struct {
	DB *gorm.DB
}

func NewCashRepository(DB *gorm.DB) CashRepository {
	return &cashRepository{
		DB: DB,
	}
}

func (cash *cashRepository) CashReceipt(cashEntity entities.CashReceipt) error {
	// debet struct
	var cashModelDebet models.CashModel = models.CashModel{
		ID:          uuid.NewString(),
		UserID:      cashEntity.UserID,
		Code:        cashEntity.Journal.Debet.Code,
		Name:        cashEntity.Journal.Debet.Name,
		Amount:      cashEntity.Journal.Debet.Amount,
		Description: cashEntity.Journal.Debet.Description,
	}
	// credit struct
	var cashModelCredit models.CashModel = models.CashModel{
		ID:          uuid.NewString(),
		UserID:      cashEntity.UserID,
		Code:        cashEntity.Journal.Credit.Code,
		Name:        cashEntity.Journal.Credit.Name,
		Amount:      cashEntity.Journal.Credit.Amount,
		Description: cashEntity.Journal.Credit.Description,
	}

	// debet query
	queryDebet := fmt.Sprintf(`
		INSERT INTO finance.general_ledgers(id, coa, description, debet, user_id) VALUES ('%s','%s', '%s', %f, '%s')
	`, cashModelDebet.ID, cashModelDebet.Code, cashModelDebet.Description, cashModelDebet.Amount, cashModelDebet.UserID)
	// credit query
	queryCredit := fmt.Sprintf(`
		INSERT INTO finance.general_ledgers (id, coa, description, credit, user_id) VALUES ('%s', '%s', '%s', '%f', '%s')
	`, cashModelCredit.ID, cashModelCredit.Code, cashModelCredit.Description, cashModelCredit.Amount, cashModelCredit.UserID)

	tx := cash.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// insert debet
	if err := tx.Exec(queryDebet).Error; err != nil {
		tx.Rollback()
		return err
	}
	// insert credit
	if err := tx.Exec(queryCredit).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
