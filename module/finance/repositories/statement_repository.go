package repositories

import (
	"github.com/riyan-eng/api-finance/module/finance/repositories/models"
	"github.com/riyan-eng/api-finance/module/finance/services/entities"
	"gorm.io/gorm"
)

type StatementRepository interface {
	IncomeStatement() (entities.LabaRugi, error)
	CapitalStatement() (entities.PerubahanModal, error)
	BalanceSheet() error
	CashFlow() error
}

type databaseStatementRepository struct {
	DB *gorm.DB
}

func NewStatementRepository(db *gorm.DB) StatementRepository {
	return &databaseStatementRepository{
		DB: db,
	}
}

func (database *databaseStatementRepository) IncomeStatement() (entities.LabaRugi, error) {
	modelLabaRugi := models.LabaRugi{
		Penjualan:                100000000,
		ReturPenjualan:           8000000,
		PotonganPenjualan:        12000000,
		PersediaanAwal:           18000000,
		Pembelian:                34000000,
		BebanAngkutPembelian:     13000000,
		ReturPembelian:           4500000,
		PotonganPembelian:        13000000,
		PersediaanAkhir:          12000000,
		BebanOperasional:         37000000,
		PendapatanNonOperasional: 2100000,
		BebanNonOperasional:      3400000,
	}

	entityPenjualan := entities.Penjualan{
		Penjualan:         modelLabaRugi.Penjualan,
		ReturPenjualan:    modelLabaRugi.ReturPenjualan,
		PotonganPenjualan: modelLabaRugi.PotonganPenjualan,
		PenjualanBersih:   modelLabaRugi.Penjualan - modelLabaRugi.ReturPenjualan - modelLabaRugi.PotonganPenjualan,
	}

	entityPembelian := entities.Pembelian{
		Pembelian:            modelLabaRugi.Pembelian,
		BebanAngkutPembelian: modelLabaRugi.BebanAngkutPembelian,
		ReturPembelian:       modelLabaRugi.ReturPembelian,
		PotonganPembelian:    modelLabaRugi.PotonganPembelian,
		PembelianBersih:      modelLabaRugi.Pembelian + modelLabaRugi.BebanAngkutPembelian - modelLabaRugi.ReturPembelian - modelLabaRugi.PotonganPembelian,
	}

	entityHpp := entities.HPP{
		PersediaanAwal:  modelLabaRugi.PersediaanAwal,
		Pembelian:       entityPembelian,
		BarangTersedia:  modelLabaRugi.PersediaanAwal + entityPembelian.PembelianBersih,
		PersediaanAkhir: modelLabaRugi.PersediaanAkhir,
		HPP:             modelLabaRugi.PersediaanAwal + entityPembelian.PembelianBersih - modelLabaRugi.PersediaanAkhir,
	}

	entityLabaRugiOperasional := entities.LabaRugiOperasional{
		Penjualan:           entityPenjualan,
		HPP:                 entityHpp,
		BebanOperasional:    modelLabaRugi.BebanOperasional,
		LabaRugiOperasional: entityPenjualan.PenjualanBersih - entityHpp.HPP - modelLabaRugi.BebanOperasional,
	}

	entityLabaRugiNonOperasional := entities.LabaRugiNonOperasional{
		PendapatanNonOperasional: modelLabaRugi.PendapatanNonOperasional,
		BebanNonOperasional:      modelLabaRugi.BebanNonOperasional,
		LabaRugiNonOperasional:   modelLabaRugi.PendapatanNonOperasional - modelLabaRugi.BebanNonOperasional,
	}

	pph := (entityLabaRugiOperasional.LabaRugiOperasional + entityLabaRugiNonOperasional.LabaRugiNonOperasional) * 5 / 100

	entityLabaRugi := entities.LabaRugi{
		LabaRugiOperasional:    entityLabaRugiOperasional,
		LabaRugiNonOperasional: entityLabaRugiNonOperasional,
		LabaBersihSebelumPajak: entityLabaRugiOperasional.LabaRugiOperasional + entityLabaRugiNonOperasional.LabaRugiNonOperasional,
		PajakPenghasilan:       pph,
		LabaBersihSetelahPajak: (entityLabaRugiOperasional.LabaRugiOperasional + entityLabaRugiNonOperasional.LabaRugiNonOperasional) - pph,
	}
	return entityLabaRugi, nil
}

func (database *databaseStatementRepository) CapitalStatement() (entities.PerubahanModal, error) {
	modelPerubahanModal := models.PerubahanModal{
		ModalAwal: 30000000,
		Prive:     1200000,
	}

	entityPerubahanModal := entities.PerubahanModal{
		ModalAwal: modelPerubahanModal.ModalAwal,
		Prive:     modelPerubahanModal.Prive,
	}
	return entityPerubahanModal, nil
}

func (database *databaseStatementRepository) BalanceSheet() error {
	return nil
}

func (database *databaseStatementRepository) CashFlow() error {
	return nil
}
