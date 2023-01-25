package repositories

import "gorm.io/gorm"

type NeracaRepository interface {
	TrialBalance()
	TrialBalanceAfterAdjustment()
	BalanceSheet()
}

type database struct {
	DB *gorm.DB
}

func NewNeracaRepository(DB *gorm.DB) NeracaRepository {
	return &database{
		DB: DB,
	}
}

func (database database) TrialBalance() {

}

func (database database) TrialBalanceAfterAdjustment() {

}

func (database database) BalanceSheet() {

}
