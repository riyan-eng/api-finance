package entities

type CashReceipt struct {
	ID      string
	UserID  string
	Journal Journal
}

type Journal struct {
	Debet  Transaction
	Credit Transaction
}

type Transaction struct {
	Code        string
	Name        string
	Amount      float64
	Description string
}
