package dto

type TrialBalanceReq struct {
	StartDate string `validate:"required" query:"start_date"`
	EndDate   string `validate:"required" query:"end_date"`
}
