package services

import (
	"github.com/riyan-eng/api-finance/config/constant"
	"github.com/riyan-eng/api-finance/src/finance/controllers/dto"
	"github.com/riyan-eng/api-finance/src/finance/repositories"
	"github.com/riyan-eng/api-finance/src/finance/services/entities"
)

type CashService interface {
	CashReceipt(data dto.CashReceiptReq) error
}

type cashService struct {
	CashRepository repositories.CashRepository
}

func NewCashService(cashRepository repositories.CashRepository) CashService {
	return &cashService{
		CashRepository: cashRepository,
	}
}

func (repo cashService) CashReceipt(data dto.CashReceiptReq) error {
	journal := entities.CashReceipt{
		ID:     "",
		UserID: "58dd4ecc-8cde-4ca3-b4f0-7451b7b59ce8",
		Journal: entities.Journal{
			Debet: entities.Transaction{
				Code:        constant.COA_CASH,
				Name:        "",
				Amount:      data.Amount,
				Description: data.Description,
			},
			Credit: entities.Transaction{
				Code:        constant.SALES,
				Name:        "",
				Amount:      data.Amount,
				Description: data.Description,
			},
		},
	}

	err := repo.CashRepository.CashReceipt(journal)
	if err != nil {
		return err
	}
	return nil
}
