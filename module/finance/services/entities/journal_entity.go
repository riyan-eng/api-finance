package entities

type JournalEntity struct {
	ID          string
	UserID      string
	Position    Position
	Amount      float64
	Description string
}

type Position struct {
	Debet  Transaction
	Credit Transaction
}

type Transaction struct {
	Code   string
	Name   string
	Amount float64
}
