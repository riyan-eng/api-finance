package services

type NeracaService interface {
	TrialBalance()
	TrialBalanceAfterAdjustment()
	BalanceSheet()
}
