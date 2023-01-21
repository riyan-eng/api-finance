package database

import "gorm.io/gorm"

type COA struct {
	gorm.Model
	ID     string `gorm:"primary_key"`
	Code   string
	Name   string
	Parent string
}

type GeneralLedger struct {
	gorm.Model
	ID          string `gorm:"primary_key"`
	COA         string
	COAID       COA `gorm:"foreignKey:COA; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Description string
	Debet       float64
	Credit      float64
	UserID      string
}
