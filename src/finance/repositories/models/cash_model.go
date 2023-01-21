package models

import "time"

type CashModel struct {
	ID        string
	CreatedAt time.Time
	Code      string
	Name      string
	Parent    string
}
