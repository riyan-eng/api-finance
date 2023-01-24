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
	CashPayment(cashEntity entities.CashReceipt) error
	Sales(cashEntity entities.CashReceipt) error
	Purchase(cashEntity entities.CashReceipt) error
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
		ID:     uuid.NewString(),
		UserID: cashEntity.UserID,
		Code:   cashEntity.Journal.Debet.Code,
		Name:   cashEntity.Journal.Debet.Name,
		Amount: cashEntity.Journal.Debet.Amount,
	}
	// credit struct
	var cashModelCredit models.CashModel = models.CashModel{
		ID:     uuid.NewString(),
		UserID: cashEntity.UserID,
		Code:   cashEntity.Journal.Credit.Code,
		Name:   cashEntity.Journal.Credit.Name,
		Amount: cashEntity.Journal.Credit.Amount,
	}

	// transaction query
	queryTransaction := fmt.Sprintf(`
		INSERT INTO finance.transactions(id, date_time, description, amount, user_id) VALUES('%v', current_timestamp,'%v', '%f', '%v')
	`, cashEntity.ID, cashEntity.Description, cashEntity.Amount, cashEntity.UserID)
	// debet query
	queryDebet := fmt.Sprintf(`
		INSERT INTO finance.general_ledgers(id, transaction, coa, debet, user_id) VALUES ('%s','%s', '%s', %f, '%s')
	`, cashModelDebet.ID, cashEntity.ID, cashModelDebet.Code, cashModelDebet.Amount, cashModelDebet.UserID)
	// credit query
	queryCredit := fmt.Sprintf(`
		INSERT INTO finance.general_ledgers (id, transaction, coa, credit, user_id) VALUES ('%s', '%s', '%s', '%f', '%s')
	`, cashModelCredit.ID, cashEntity.ID, cashModelCredit.Code, cashModelCredit.Amount, cashModelCredit.UserID)

	tx := cash.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// insert transaction
	if err := tx.Exec(queryTransaction).Error; err != nil {
		tx.Rollback()
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

func (cash *cashRepository) CashPayment(cashEntity entities.CashReceipt) error {
	// debet struct
	var cashModelDebet models.CashModel = models.CashModel{
		ID:     uuid.NewString(),
		UserID: cashEntity.UserID,
		Code:   cashEntity.Journal.Debet.Code,
		Name:   cashEntity.Journal.Debet.Name,
		Amount: cashEntity.Journal.Debet.Amount,
	}
	// credit struct
	var cashModelCredit models.CashModel = models.CashModel{
		ID:     uuid.NewString(),
		UserID: cashEntity.UserID,
		Code:   cashEntity.Journal.Credit.Code,
		Name:   cashEntity.Journal.Credit.Name,
		Amount: cashEntity.Journal.Credit.Amount,
	}

	// transaction query
	queryTransaction := fmt.Sprintf(`
		INSERT INTO finance.transactions (id, date_time, description, amount, user_id) VALUES ('%s', current_timestamp, '%s', '%f', '%s')
	`, cashEntity.ID, cashEntity.Description, cashEntity.Amount, cashEntity.UserID)
	// debet query
	queryDebet := fmt.Sprintf(`
		INSERT INTO finance.general_ledgers (id, transaction, coa, debet, user_id) VALUES ('%s', '%s', '%s', '%f', '%s')
	`, cashModelDebet.ID, cashEntity.ID, cashModelDebet.Code, cashModelDebet.Amount, cashModelDebet.UserID)
	// credit query
	queryCredit := fmt.Sprintf(`
		INSERT INTO finance.general_ledgers (id, transaction, coa, credit, user_id) VALUES ('%s', '%s', '%s', '%f', '%s')
	`, cashModelCredit.ID, cashEntity.ID, cashModelCredit.Code, cashModelCredit.Amount, cashModelCredit.UserID)

	tx := cash.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// insert transaction
	if err := tx.Exec(queryTransaction).Error; err != nil {
		tx.Rollback()
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

func (cash *cashRepository) Sales(cashEntity entities.CashReceipt) error {
	// debet struct
	var cashModelDebet models.CashModel = models.CashModel{
		ID:     uuid.NewString(),
		UserID: cashEntity.UserID,
		Code:   cashEntity.Journal.Debet.Code,
		Name:   cashEntity.Journal.Debet.Name,
		Amount: cashEntity.Journal.Debet.Amount,
	}
	// credit struct
	var cashModelCredit models.CashModel = models.CashModel{
		ID:     uuid.NewString(),
		UserID: cashEntity.UserID,
		Code:   cashEntity.Journal.Credit.Code,
		Name:   cashEntity.Journal.Credit.Name,
		Amount: cashEntity.Journal.Credit.Amount,
	}

	// transaction query
	queryTransaction := fmt.Sprintf(`
		INSERT INTO finance.transactions (id, date_time, description, amount, user_id) VALUES ('%s', current_timestamp, '%s', '%f', '%s')
	`, cashEntity.ID, cashEntity.Description, cashEntity.Amount, cashEntity.UserID)
	queryDebet := fmt.Sprintf(`
		INSERT INTO finance.general_ledgers (id, transaction, coa, debet, user_id) VALUES ('%s', '%s', '%s', '%f', '%s')
	`, cashModelDebet.ID, cashEntity.ID, cashModelDebet.Code, cashModelDebet.Amount, cashModelDebet.UserID)
	// credit query
	queryCredit := fmt.Sprintf(`
		INSERT INTO finance.general_ledgers (id, transaction, coa, credit, user_id) VALUES ('%s', '%s', '%s', '%f', '%s')
	`, cashModelCredit.ID, cashEntity.ID, cashModelCredit.Code, cashModelCredit.Amount, cashModelCredit.UserID)

	tx := cash.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// insert transaction
	if err := tx.Exec(queryTransaction).Error; err != nil {
		tx.Rollback()
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

func (cash *cashRepository) Purchase(cashEntity entities.CashReceipt) error {
	// debet struct
	var cashModelDebet models.CashModel = models.CashModel{
		ID:     uuid.NewString(),
		UserID: cashEntity.UserID,
		Code:   cashEntity.Journal.Debet.Code,
		Name:   cashEntity.Journal.Debet.Name,
		Amount: cashEntity.Journal.Debet.Amount,
	}
	// credit struct
	var cashModelCredit models.CashModel = models.CashModel{
		ID:     uuid.NewString(),
		UserID: cashEntity.UserID,
		Code:   cashEntity.Journal.Credit.Code,
		Name:   cashEntity.Journal.Credit.Name,
		Amount: cashEntity.Journal.Credit.Amount,
	}

	// transaction query
	queryTransaction := fmt.Sprintf(`
		INSERT INTO finance.transactions (id, date_time, description, amount, user_id) VALUES ('%s', current_timestamp, '%s', '%f', '%s')
	`, cashEntity.ID, cashEntity.Description, cashEntity.Amount, cashEntity.UserID)
	queryDebet := fmt.Sprintf(`
		INSERT INTO finance.general_ledgers (id, transaction, coa, debet, user_id) VALUES ('%s', '%s', '%s', '%f', '%s')
	`, cashModelDebet.ID, cashEntity.ID, cashModelDebet.Code, cashModelDebet.Amount, cashModelDebet.UserID)
	// credit query
	queryCredit := fmt.Sprintf(`
		INSERT INTO finance.general_ledgers (id, transaction, coa, credit, user_id) VALUES ('%s', '%s', '%s', '%f', '%s')
	`, cashModelCredit.ID, cashEntity.ID, cashModelCredit.Code, cashModelCredit.Amount, cashModelCredit.UserID)

	tx := cash.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// insert transaction
	if err := tx.Exec(queryTransaction).Error; err != nil {
		tx.Rollback()
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
