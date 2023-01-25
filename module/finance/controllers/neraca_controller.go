package controllers

type NeracaController interface {
	TrialBalance()
	TrialBalanceAfterAdjustment()
	BalanceSheet()
}
