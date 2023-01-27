package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/api-finance/module/finance/controllers/dto"
	"github.com/riyan-eng/api-finance/module/finance/services"
)

type StatementController interface {
	IncomeStatement(c *fiber.Ctx) error
	CapitalStatement(c *fiber.Ctx) error
	BalanceSheet(c *fiber.Ctx) error
	CashFlow(c *fiber.Ctx) error
}

type statementService struct {
	Service services.StatementService
}

func NewStatementController(sS services.StatementService) StatementController {
	return &statementService{
		Service: sS,
	}
}

func (service *statementService) IncomeStatement(c *fiber.Ctx) error {
	entityIncomeStatement, err := service.Service.IncomeStatement()
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	incomeStatementRes := dto.LabaRugiRes{
		LabaRugiOperasional: dto.LabaRugiOperasional{
			Penjualan: dto.Penjualan{
				Penjualan:         entityIncomeStatement.LabaRugiOperasional.Penjualan.Penjualan,
				ReturPenjualan:    entityIncomeStatement.LabaRugiOperasional.Penjualan.ReturPenjualan,
				PotonganPenjualan: entityIncomeStatement.LabaRugiOperasional.Penjualan.PotonganPenjualan,
				PenjualanBersih:   entityIncomeStatement.LabaRugiOperasional.Penjualan.PenjualanBersih,
			},
			HPP: dto.HPP{
				PersediaanAwal: entityIncomeStatement.LabaRugiOperasional.HPP.PersediaanAwal,
				Pembelian: dto.Pembelian{
					Pembelian:            entityIncomeStatement.LabaRugiOperasional.HPP.Pembelian.Pembelian,
					BebanAngkutPembelian: entityIncomeStatement.LabaRugiOperasional.HPP.Pembelian.BebanAngkutPembelian,
					ReturPembelian:       entityIncomeStatement.LabaRugiOperasional.HPP.Pembelian.ReturPembelian,
					PotonganPembelian:    entityIncomeStatement.LabaRugiOperasional.HPP.Pembelian.PotonganPembelian,
					PembelianBersih:      entityIncomeStatement.LabaRugiOperasional.HPP.Pembelian.PembelianBersih,
				},
				BarangTersedia:  entityIncomeStatement.LabaRugiOperasional.HPP.BarangTersedia,
				PersediaanAkhir: entityIncomeStatement.LabaRugiOperasional.HPP.PersediaanAkhir,
				HPP:             entityIncomeStatement.LabaRugiOperasional.HPP.HPP,
			},
			BebanOperasional:    entityIncomeStatement.LabaRugiOperasional.BebanOperasional,
			LabaRugiOperasional: entityIncomeStatement.LabaRugiOperasional.LabaRugiOperasional,
		},
		LabaRugiNonOperasional: dto.LabaRugiNonOperasional{
			PendapatanNonOperasional: entityIncomeStatement.LabaRugiNonOperasional.PendapatanNonOperasional,
			BebanNonOperasional:      entityIncomeStatement.LabaRugiNonOperasional.BebanNonOperasional,
			LabaRugiNonOperasional:   entityIncomeStatement.LabaRugiNonOperasional.LabaRugiNonOperasional,
		},
		LabaBersihSebelumPajak: entityIncomeStatement.LabaBersihSebelumPajak,
		PajakPenghasilan:       entityIncomeStatement.PajakPenghasilan,
		LabaBersihSetelahPajak: entityIncomeStatement.LabaBersihSetelahPajak,
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    incomeStatementRes,
		"message": "ok",
	})
}

func (service *statementService) CapitalStatement(c *fiber.Ctx) error {
	entityPerubahanModal, err := service.Service.CapitalStatement()
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	perubahanModalRes := dto.PerubahanModalRes{
		ModalAwal:  entityPerubahanModal.ModalAwal,
		LabaRugi:   entityPerubahanModal.LabaRugi,
		Prive:      entityPerubahanModal.Prive,
		ModalAkhir: entityPerubahanModal.ModalAkhir,
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    perubahanModalRes,
		"message": "ok",
	})
}

func (service *statementService) BalanceSheet(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    1,
		"message": "ok",
	})
}

func (service *statementService) CashFlow(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    1,
		"message": "ok",
	})
}
