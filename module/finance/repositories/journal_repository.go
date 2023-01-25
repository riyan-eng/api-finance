package repositories

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/riyan-eng/api-finance/module/finance/repositories/models"
	"github.com/riyan-eng/api-finance/module/finance/services/entities"
	"gorm.io/gorm"
)

type JournalRepository interface {
	CashReceipt(journal entities.JournalEntity) error
	CashPayment(journal entities.JournalEntity) error
	Sales(journal entities.JournalEntity) error
	Purchase(journal entities.JournalEntity) error
	General(journal entities.JournalEntity) error
}

type journalRepository struct {
	DB *gorm.DB
}

func NewJournalRepository(DB *gorm.DB) JournalRepository {
	return &journalRepository{
		DB: DB,
	}
}

func (jR *journalRepository) CashReceipt(journal entities.JournalEntity) error {
	// debet struct
	var jurnalDebet models.JournalModel = models.JournalModel{
		ID:     uuid.NewString(),
		UserID: journal.UserID,
		Code:   journal.Position.Debet.Code,
		Name:   journal.Position.Debet.Name,
		Amount: journal.Position.Debet.Amount,
	}
	// credit struct
	var journalCredit models.JournalModel = models.JournalModel{
		ID:     uuid.NewString(),
		UserID: journal.UserID,
		Code:   journal.Position.Credit.Code,
		Name:   journal.Position.Credit.Name,
		Amount: journal.Position.Credit.Amount,
	}

	// transaction query
	queryTransaction := fmt.Sprintf(`
		INSERT INTO finance.transactions(id, date_time, description, amount, user_id) VALUES('%v', current_timestamp,'%v', '%f', '%v')
	`, journal.ID, journal.Description, journal.Amount, journal.UserID)
	// debet query
	queryDebet := fmt.Sprintf(`
		INSERT INTO finance.cash_receipt_journals(id, transaction, coa, debet, user_id) VALUES ('%s','%s', '%s', %f, '%s')
	`, jurnalDebet.ID, journal.ID, jurnalDebet.Code, jurnalDebet.Amount, jurnalDebet.UserID)
	// credit query
	queryCredit := fmt.Sprintf(`
		INSERT INTO finance.cash_receipt_journals (id, transaction, coa, credit, user_id) VALUES ('%s', '%s', '%s', '%f', '%s')
	`, journalCredit.ID, journal.ID, journalCredit.Code, journalCredit.Amount, journalCredit.UserID)

	tx := jR.DB.Begin()
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

func (jR *journalRepository) CashPayment(journal entities.JournalEntity) error {
	// debet struct
	var jurnalDebet models.JournalModel = models.JournalModel{
		ID:     uuid.NewString(),
		UserID: journal.UserID,
		Code:   journal.Position.Debet.Code,
		Name:   journal.Position.Debet.Name,
		Amount: journal.Position.Debet.Amount,
	}
	// credit struct
	var journalCredit models.JournalModel = models.JournalModel{
		ID:     uuid.NewString(),
		UserID: journal.UserID,
		Code:   journal.Position.Credit.Code,
		Name:   journal.Position.Credit.Name,
		Amount: journal.Position.Credit.Amount,
	}

	// transaction query
	queryTransaction := fmt.Sprintf(`
		INSERT INTO finance.transactions (id, date_time, description, amount, user_id) VALUES ('%s', current_timestamp, '%s', '%f', '%s')
	`, journal.ID, journal.Description, journal.Amount, journal.UserID)
	// debet query
	queryDebet := fmt.Sprintf(`
		INSERT INTO finance.cash_payment_journals (id, transaction, coa, debet, user_id) VALUES ('%s', '%s', '%s', '%f', '%s')
	`, jurnalDebet.ID, journal.ID, jurnalDebet.Code, jurnalDebet.Amount, jurnalDebet.UserID)
	// credit query
	queryCredit := fmt.Sprintf(`
		INSERT INTO finance.cash_payment_journals (id, transaction, coa, credit, user_id) VALUES ('%s', '%s', '%s', '%f', '%s')
	`, journalCredit.ID, journal.ID, journalCredit.Code, journalCredit.Amount, journalCredit.UserID)

	tx := jR.DB.Begin()
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

func (jR *journalRepository) Sales(journal entities.JournalEntity) error {
	// debet struct
	var jurnalDebet models.JournalModel = models.JournalModel{
		ID:     uuid.NewString(),
		UserID: journal.UserID,
		Code:   journal.Position.Debet.Code,
		Name:   journal.Position.Debet.Name,
		Amount: journal.Position.Debet.Amount,
	}
	// credit struct
	var journalCredit models.JournalModel = models.JournalModel{
		ID:     uuid.NewString(),
		UserID: journal.UserID,
		Code:   journal.Position.Credit.Code,
		Name:   journal.Position.Credit.Name,
		Amount: journal.Position.Credit.Amount,
	}

	// transaction query
	queryTransaction := fmt.Sprintf(`
		INSERT INTO finance.transactions (id, date_time, description, amount, user_id) VALUES ('%s', current_timestamp, '%s', '%f', '%s')
	`, journal.ID, journal.Description, journal.Amount, journal.UserID)
	queryDebet := fmt.Sprintf(`
		INSERT INTO finance.sales_journals (id, transaction, coa, debet, user_id) VALUES ('%s', '%s', '%s', '%f', '%s')
	`, jurnalDebet.ID, journal.ID, jurnalDebet.Code, jurnalDebet.Amount, jurnalDebet.UserID)
	// credit query
	queryCredit := fmt.Sprintf(`
		INSERT INTO finance.sales_journals (id, transaction, coa, credit, user_id) VALUES ('%s', '%s', '%s', '%f', '%s')
	`, journalCredit.ID, journal.ID, journalCredit.Code, journalCredit.Amount, journalCredit.UserID)

	tx := jR.DB.Begin()
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

func (jR *journalRepository) Purchase(journal entities.JournalEntity) error {
	// debet struct
	var jurnalDebet models.JournalModel = models.JournalModel{
		ID:     uuid.NewString(),
		UserID: journal.UserID,
		Code:   journal.Position.Debet.Code,
		Name:   journal.Position.Debet.Name,
		Amount: journal.Position.Debet.Amount,
	}
	// credit struct
	var journalCredit models.JournalModel = models.JournalModel{
		ID:     uuid.NewString(),
		UserID: journal.UserID,
		Code:   journal.Position.Credit.Code,
		Name:   journal.Position.Credit.Name,
		Amount: journal.Position.Credit.Amount,
	}

	// transaction query
	queryTransaction := fmt.Sprintf(`
		INSERT INTO finance.transactions (id, date_time, description, amount, user_id) VALUES ('%s', current_timestamp, '%s', '%f', '%s')
	`, journal.ID, journal.Description, journal.Amount, journal.UserID)
	queryDebet := fmt.Sprintf(`
		INSERT INTO finance.purchase_journals (id, transaction, coa, debet, user_id) VALUES ('%s', '%s', '%s', '%f', '%s')
	`, jurnalDebet.ID, journal.ID, jurnalDebet.Code, jurnalDebet.Amount, jurnalDebet.UserID)
	// credit query
	queryCredit := fmt.Sprintf(`
		INSERT INTO finance.purchase_journals (id, transaction, coa, credit, user_id) VALUES ('%s', '%s', '%s', '%f', '%s')
	`, journalCredit.ID, journal.ID, journalCredit.Code, journalCredit.Amount, journalCredit.UserID)

	tx := jR.DB.Begin()
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

func (jR *journalRepository) General(journal entities.JournalEntity) error {
	// debet struct
	var jurnalDebet models.JournalModel = models.JournalModel{
		ID:     uuid.NewString(),
		UserID: journal.UserID,
		Code:   journal.Position.Debet.Code,
		Name:   journal.Position.Debet.Name,
		Amount: journal.Position.Debet.Amount,
	}
	// credit struct
	var journalCredit models.JournalModel = models.JournalModel{
		ID:     uuid.NewString(),
		UserID: journal.UserID,
		Code:   journal.Position.Credit.Code,
		Name:   journal.Position.Credit.Name,
		Amount: journal.Position.Credit.Amount,
	}

	// transaction query
	queryTransaction := fmt.Sprintf(`
		INSERT INTO finance.transactions (id, date_time, description, amount, user_id) VALUES ('%s', current_timestamp, '%s', '%f', '%s')
	`, journal.ID, journal.Description, journal.Amount, journal.UserID)
	queryDebet := fmt.Sprintf(`
		INSERT INTO finance.general_ledgers (id, transaction, coa, debet, user_id) VALUES ('%s', '%s', '%s', '%f', '%s')
	`, jurnalDebet.ID, journal.ID, jurnalDebet.Code, jurnalDebet.Amount, jurnalDebet.UserID)
	// credit query
	queryCredit := fmt.Sprintf(`
		INSERT INTO finance.general_ledgers (id, transaction, coa, credit, user_id) VALUES ('%s', '%s', '%s', '%f', '%s')
	`, journalCredit.ID, journal.ID, journalCredit.Code, journalCredit.Amount, journalCredit.UserID)

	tx := jR.DB.Begin()
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
