package util

import (
	"time"

	"gorm.io/gorm"
)

type COA struct {
	gorm.Model
	ID     string `gorm:"primary_key"`
	Code   string `gorm:"unique"`
	Name   string
	Parent string
}

type Transaction struct {
	gorm.Model
	ID          string `gorm:"primary_key"`
	DateTime    time.Time
	Description string
	Amount      float64
	UserID      string
}

type GeneralLedger struct {
	gorm.Model
	ID            string `gorm:"primary_key"`
	Transaction   string
	TransactionID Transaction `gorm:"foreignKey:Transaction; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	COA           string
	COAID         COA `gorm:"foreignKey:COA; references:Code; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Debet         float64
	Credit        float64
	UserID        string
}

type CashReceiptJournal struct {
	gorm.Model
	ID            string `gorm:"primary_key"`
	Transaction   string
	TransactionID Transaction `gorm:"foreignKey:Transaction; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	COA           string
	COAID         COA `gorm:"foreignKey:COA; references:Code; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Debet         float64
	Credit        float64
	UserID        string
}

type CashPaymentJournal struct {
	gorm.Model
	ID            string `gorm:"primary_key"`
	Transaction   string
	TransactionID Transaction `gorm:"foreignKey:Transaction; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	COA           string
	COAID         COA `gorm:"foreignKey:COA; references:Code; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Debet         float64
	Credit        float64
	UserID        string
}

type SalesJournal struct {
	gorm.Model
	ID            string `gorm:"primary_key"`
	Transaction   string
	TransactionID Transaction `gorm:"foreignKey:Transaction; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	COA           string
	COAID         COA `gorm:"foreignKey:COA; references:Code; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Debet         float64
	Credit        float64
	UserID        string
}

type PurchaseJournal struct {
	gorm.Model
	ID            string `gorm:"primary_key"`
	Transaction   string
	TransactionID Transaction `gorm:"foreignKey:Transaction; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	COA           string
	COAID         COA `gorm:"foreignKey:COA; references:Code; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Debet         float64
	Credit        float64
	UserID        string
}

type AdjustmentJournalEntry struct {
	gorm.Model
	ID            string `gorm:"primary_key"`
	Transaction   string
	TransactionID Transaction `gorm:"foreignKey:Transaction; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	COA           string
	COAID         COA `gorm:"foreignKey:COA; references:Code; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Debet         float64
	Credit        float64
	UserID        string
}
