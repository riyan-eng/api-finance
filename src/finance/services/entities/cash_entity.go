package entities

type CashReceipt struct {
	ID      string
	UserId  string
	Journal journal
}

type journal struct {
	Debet  transaction
	Credit transaction
}

type transaction struct {
	Code   string
	Name   string
	Amount float64
}
