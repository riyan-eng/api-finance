package dto

type CashReceiptReq struct {
	Code        string  `validate:"required" json:"code"`
	Description string  `validate:"required" json:"description"`
	Amount      float64 `validate:"required" json:"amount"`
}
