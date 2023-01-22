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
	var err error

	// debet
	var cashModelDebet models.CashModel = models.CashModel{
		ID:          uuid.NewString(),
		UserID:      cashEntity.UserID,
		Code:        cashEntity.Journal.Debet.Code,
		Name:        cashEntity.Journal.Debet.Name,
		Amount:      cashEntity.Journal.Debet.Amount,
		Description: cashEntity.Journal.Debet.Description,
	}
	queryDebet := fmt.Sprintf(`
		INSERT INTO finance.general_ledgers (id, coa, description, debet, user_id) VALUES (%v, %v, %v, %v, %v)
	`, cashModelDebet.ID, cashModelDebet.Code, cashModelDebet.Description, cashModelDebet.Amount, cashModelDebet.UserID)
	err = cash.DB.Exec(queryDebet).Error
	fmt.Println("--- repository ---")
	fmt.Println(err)
	fmt.Println("--- repository ---")
	if err != nil {
		return err
	}

	// credit
	var cashModelCredit models.CashModel = models.CashModel{
		ID:          uuid.NewString(),
		UserID:      cashEntity.UserID,
		Code:        cashEntity.Journal.Debet.Code,
		Name:        cashEntity.Journal.Debet.Name,
		Amount:      cashEntity.Journal.Debet.Amount,
		Description: cashEntity.Journal.Debet.Description,
	}
	queryCredit := fmt.Sprintf(`
		INSERT INTO finance.general_ledgers (id, coa, description, credit, user_id) VALUES (%v, %v, %v, %v, %v)
	`, cashModelCredit.ID, cashModelCredit.Code, cashModelCredit.Description, cashModelCredit.Amount, cashModelCredit.UserID)
	err = cash.DB.Exec(queryCredit).Error
	fmt.Println("--- repository ---")
	fmt.Println(err)
	fmt.Println("--- repository ---")
	if err != nil {
		return err
	}

	return nil
}
