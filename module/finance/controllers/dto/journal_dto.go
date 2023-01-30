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
	GoodCode string `validate:"required" json:"good_code"`
	Quantity int    `validate:"required" json:"quantity"`
	TaxCode  string `json:"tax_code"`
}

type SalesRes struct {
	GoodCode string
	Quantity int
	Amount   float64
	Tax      float64
	Total    float64
}

type PurchaseReq struct {
	Code        string  `validate:"required" json:"code"`
	Description string  `validate:"required" json:"description"`
	Amount      float64 `validate:"required" json:"amount"`
}

type GeneralReq struct {
	CodeDebet   string  `validate:"required" json:"code_debet"`
	CodeCredit  string  `validate:"required" json:"code_credit"`
	Description string  `validate:"required" json:"description"`
	Amount      float64 `validate:"required" json:"amount"`
}
