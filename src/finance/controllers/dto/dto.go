package dto

type CashReceiptReq struct {
	Amount      float64 `validate:"required" json:"amount"`
	Description string  `validate:"required" json:"description"`
}
