package repositories

import "gorm.io/gorm"

type StatementRepository interface {
	IncomeStatement() error
	CapitalStatement() error
	BalanceSheet() error
	CashFlow() error
}

type databaseStatementRepository struct {
	DB *gorm.DB
}

func NewStatementRepository(db *gorm.DB) StatementRepository {
	return &databaseStatementRepository{
		DB: db,
	}
}

func (database *databaseStatementRepository) IncomeStatement() error {
	return nil
}

func (database *databaseStatementRepository) CapitalStatement() error {
	return nil
}

func (database *databaseStatementRepository) BalanceSheet() error {
	return nil
}

func (database *databaseStatementRepository) CashFlow() error {
	return nil
}
