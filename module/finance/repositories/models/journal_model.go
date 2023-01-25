package models

import "time"

type JournalModel struct {
	ID          string
	UserID      string
	CreatedAt   time.Time
	Code        string
	Name        string
	Amount      float64
	Description string
}
