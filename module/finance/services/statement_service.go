package services

import "github.com/riyan-eng/api-finance/module/finance/repositories"

type StatementService interface {
	IncomeStatement()
	CapitalStatement()
	BalanceSheet()
	CashFlow()
}

type statementRepository struct {
	Repo repositories.StatementRepository
}

func NewStatementService(sR repositories.StatementRepository) StatementService {
	return &statementRepository{
		Repo: sR,
	}
}

func (repo *statementRepository) IncomeStatement() {

}

func (repo *statementRepository) CapitalStatement() {

}

func (repo *statementRepository) BalanceSheet() {

}

func (repo *statementRepository) CashFlow() {

}
