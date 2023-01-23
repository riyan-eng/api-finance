package dto

type CashReceiptReq struct {
	Code        string  `validate:"required" json:"code"`
	Description string  `validate:"required" json:"description"`
	Amount      float64 `validate:"required" json:"amount"`
}

type CashPaymentReq struct {
	Code        string  `validate:"required" json:"code"`
	Description string  `validate:"required" json:"description"`
	Amount      float64 `validate:"required" json:"amount"`
}

type SalesReq struct {
	Code        string  `validate:"required" json:"code"`
	Description string  `validate:"required" json:"description"`
	Amount      float64 `validate:"required" json:"amount"`
}

type PurchaseReq struct {
	Code        string  `validate:"required" json:"code"`
	Description string  `validate:"required" json:"description"`
	Amount      float64 `validate:"required" json:"amount"`
}

type GeneralReq struct {
	CodeDebet   string `validate:"required" json:"code_debet"`
	CodeCredit  string `validate:"required" json:"code_credit"`
	Description string `validate:"required" json:"description"`
	Amount      string `validate:"required" json:"amount"`
}
