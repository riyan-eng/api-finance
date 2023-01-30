package services

import (
	"fmt"

	"github.com/riyan-eng/api-finance/module/finance/controllers/dto"
	"github.com/riyan-eng/api-finance/module/finance/repositories"
	"github.com/riyan-eng/api-finance/module/finance/services/entities"
)

type JournalService interface {
	CashReceipt(journal entities.JournalEntity) error
	CashPayment(journal entities.JournalEntity) error
	Sales(sR dto.SalesReq) (dto.SalesRes, error)
	Purchase(journal entities.JournalEntity) error
	General(journal entities.JournalEntity) error
}

type journalRepository struct {
	CashRepository repositories.JournalRepository
}

func NewJournalService(jR repositories.JournalRepository) JournalService {
	return &journalRepository{
		CashRepository: jR,
	}
}

func (repo *journalRepository) CashReceipt(journal entities.JournalEntity) error {
	if err := repo.CashRepository.CashReceipt(journal); err != nil {
		return err
	}
	return nil
}

func (repo *journalRepository) CashPayment(journal entities.JournalEntity) error {
	if err := repo.CashRepository.CashPayment(journal); err != nil {
		return err
	}
	return nil
}

func (repo *journalRepository) Sales(sR dto.SalesReq) (dto.SalesRes, error) {
	// fmt.Println(sR)
	var quantity = 8
	var price float64 = 2100000
	var tax_rate float64 = 11
	var tax float64

	var amount = float64(quantity) * price
	fmt.Println(amount)

	fmt.Println(sR.TaxCode)
	if sR.TaxCode != "" {
		fmt.Println("pajak")
		tax = (amount * tax_rate) / 100
	}

	fmt.Println(tax)

	var total = amount + tax
	fmt.Println(total)

	// if err := repo.CashRepository.Sales(journal); err != nil {
	// 	return err
	// }
	data := dto.SalesRes{
		GoodCode: sR.GoodCode,
		Quantity: quantity,
		Amount:   amount,
		Tax:      tax,
		Total:    total,
	}
	return data, nil
}

func (repo *journalRepository) Purchase(journal entities.JournalEntity) error {
	if err := repo.CashRepository.Purchase(journal); err != nil {
		return err
	}
	return nil
}

func (repo *journalRepository) General(journal entities.JournalEntity) error {
	if err := repo.CashRepository.General(journal); err != nil {
		return err
	}
	return nil
}
