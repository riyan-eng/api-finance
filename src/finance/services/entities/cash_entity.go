package entities

type CashReceipt struct {
	ID      string
	UserID  string
	Journal journal
}

type journal struct {
	Debet  transaction
	Credit transaction
}

type transaction struct {
	Code        string
	Name        string
	Amount      float64
	Description string
}
