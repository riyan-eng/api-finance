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

type database struct {
	DB *gorm.DB
}

func NewNeracaRepository(DB *gorm.DB) NeracaRepository {
	return &database{
		DB: DB,
	}
}

func (db database) TrialBalance(nE entities.NeracaEntity) error {
	return nil
}

func (db database) TrialBalanceAfterAdjustment(nE entities.NeracaEntity) error {
	return nil
}

func (db database) BalanceSheet(nE entities.NeracaEntity) error {
	return nil
}
