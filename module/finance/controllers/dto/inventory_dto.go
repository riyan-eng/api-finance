package dto

type IncreaseInventoryReq struct {
	Code  string  `validate:"required" json:"code"`
	Qty   uint    `validate:"required" json:"quantity"`
	Price float64 `validate:"required" json:"price"`
}

type DecreaseInventoryReq struct {
	Code string `validate:"required" json:"code"`
	Qty  uint   `validate:"required" json:"quantity"`
}
