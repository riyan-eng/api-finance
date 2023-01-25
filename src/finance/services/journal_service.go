package services

import (
	"github.com/riyan-eng/api-finance/src/finance/repositories"
	"github.com/riyan-eng/api-finance/src/finance/services/entities"
)

type JournalService interface {
	CashReceipt(journal entities.JournalEntity) error
	CashPayment(journal entities.JournalEntity) error
	Sales(journal entities.JournalEntity) error
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

func (repo *journalRepository) Sales(journal entities.JournalEntity) error {
	if err := repo.CashRepository.Sales(journal); err != nil {
		return err
	}
	return nil
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
