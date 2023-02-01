package entities

type InventoryIn struct {
	Code          string
	Price         float64
	Qty           uint
	Amount        float64
	BalancePrice  float64
	BalanceQty    uint
	BalanceAmount float64
}

type InventoryOut struct {
	Code          string
	Price         float64
	Qty           uint
	Amount        float64
	BalancePrice  float64
	BalanceQty    uint
	BalanceAmount float64
}
