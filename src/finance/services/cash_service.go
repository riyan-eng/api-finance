package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/riyan-eng/api-finance/config/constant"
	"github.com/riyan-eng/api-finance/src/finance/controllers/dto"
	"github.com/riyan-eng/api-finance/src/finance/repositories"
	"github.com/riyan-eng/api-finance/src/finance/services/entities"
)

type CashService interface {
	CashReceipt(data dto.CashReceiptReq) error
	CashPayment(data dto.CashPaymentReq) error
	Sales(data dto.SalesReq) error
}

type cashService struct {
	CashRepository repositories.CashRepository
}

func NewCashService(cashRepository repositories.CashRepository) CashService {
	return &cashService{
		CashRepository: cashRepository,
	}
}

func (repo *cashService) CashReceipt(data dto.CashReceiptReq) error {
	journal := entities.CashReceipt{
		ID:          uuid.NewString(),
		UserID:      "58dd4ecc-8cde-4ca3-b4f0-7451b7b59ce8",
		Amount:      data.Amount,
		Description: data.Description,
		Journal: entities.Journal{
			Debet: entities.Transaction{
				Code:   constant.COA_CASH,
				Amount: data.Amount,
			},
			Credit: entities.Transaction{
				Code:   data.Code,
				Amount: data.Amount,
			},
		},
	}

	err := repo.CashRepository.CashReceipt(journal)
	fmt.Println("--- service ---")
	fmt.Println(err)
	fmt.Println("--- service ---")
	if err != nil {
		return err
	}
	return nil
}

func (repo *cashService) CashPayment(data dto.CashPaymentReq) error {
	journal := entities.CashReceipt{
		ID:          uuid.NewString(),
		UserID:      "58dd4ecc-8cde-4ca3-b4f0-7451b7b59ce8",
		Amount:      data.Amount,
		Description: data.Description,
		Journal: entities.Journal{
			Debet: entities.Transaction{
				Code:   data.Code,
				Amount: data.Amount,
			},
			Credit: entities.Transaction{
				Code:   constant.COA_CASH,
				Amount: data.Amount,
			},
		},
	}
	if err := repo.CashRepository.CashPayment(journal); err != nil {
		return err
	}
	return nil
}

func (repo *cashService) Sales(data dto.SalesReq) error {
	journal := entities.CashReceipt{
		ID:          uuid.NewString(),
		UserID:      "58dd4ecc-8cde-4ca3-b4f0-7451b7b59ce8",
		Amount:      data.Amount,
		Description: data.Description,
		Journal: entities.Journal{
			Debet: entities.Transaction{
				Code:   data.Code,
				Amount: data.Amount,
			},
			Credit: entities.Transaction{
				Code:   constant.COA_SALES,
				Amount: data.Amount,
			},
		},
	}

	if err := repo.CashRepository.Sales(journal); err != nil {
		return err
	}
	return nil
}
