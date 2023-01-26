package repositories

import (
	"github.com/riyan-eng/api-finance/module/finance/services/entities"
	"gorm.io/gorm"
)

type NeracaRepository interface {
	TrialBalance(nE entities.NeracaEntity) error
	TrialBalanceAfterAdjustment(nE entities.NeracaEntity) error
	BalanceSheet(nE entities.NeracaEntity) error
}

type databaseNeracaRepository struct {
	DB *gorm.DB
}

func NewNeracaRepository(DB *gorm.DB) NeracaRepository {
	return &databaseNeracaRepository{
		DB: DB,
	}
}

func (database *databaseNeracaRepository) TrialBalance(nE entities.NeracaEntity) error {
	return nil
}

func (database *databaseNeracaRepository) TrialBalanceAfterAdjustment(nE entities.NeracaEntity) error {
	return nil
}

func (database *databaseNeracaRepository) BalanceSheet(nE entities.NeracaEntity) error {
	return nil
}
