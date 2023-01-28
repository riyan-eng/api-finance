package entities

type NeracaEntity struct {
	Code   string
	Name   string
	Debet  float64
	Credit float64
	Saldo  float64
}

type COA struct {
	Code string
	Name string
	// NameBahasa string
	Child []COAChild
}

type COAChild struct {
	Code       string
	Name       string
	NameBahasa string
}
