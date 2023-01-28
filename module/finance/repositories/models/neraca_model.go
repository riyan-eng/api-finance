package models

type NeracaModel struct {
	Code   string
	Coa    string
	Debet  float64
	Credit float64
	Saldo  float64
}

type ChartOfAccountParentModel struct {
	Code       string
	Name       string
	NameBahasa string
}

type ChartOfAccountChildModel struct {
	Parent     string
	Code       string
	Name       string
	NameBahasa string
}
