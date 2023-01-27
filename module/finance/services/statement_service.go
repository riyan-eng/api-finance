package services

import (
	"github.com/riyan-eng/api-finance/module/finance/repositories"
	"github.com/riyan-eng/api-finance/module/finance/services/entities"
)

type StatementService interface {
	IncomeStatement() (entities.LabaRugi, error)
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

func (repo *statementRepository) IncomeStatement() (entities.LabaRugi, error) {
	entityIncomeStatement, err := repo.Repo.IncomeStatement()
	if err != nil {
		return entities.LabaRugi{}, err
	}
	return entityIncomeStatement, nil
}

func (repo *statementRepository) CapitalStatement() {

}

func (repo *statementRepository) BalanceSheet() {

}

func (repo *statementRepository) CashFlow() {

}
