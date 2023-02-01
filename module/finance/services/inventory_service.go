package services

import (
	"errors"

	"github.com/riyan-eng/api-finance/module/finance/controllers/dto"
	"github.com/riyan-eng/api-finance/module/finance/repositories"
	"github.com/riyan-eng/api-finance/module/finance/services/entities"
)

type InventoryService interface {
	GetStocks() error
	IncreaseStock(*dto.IncreaseInventoryReq) error
	DecreaseStock(*dto.DecreaseInventoryReq) error
}

type repositoryInventoryService struct {
	InventoryRepository repositories.InventoryRepository
}

func NewInventoryService(iR repositories.InventoryRepository) InventoryService {
	return &repositoryInventoryService{
		InventoryRepository: iR,
	}
}

func (repo repositoryInventoryService) GetStocks() error {
	return nil
}

func (repo repositoryInventoryService) IncreaseStock(inventoryIn *dto.IncreaseInventoryReq) error {
	// get saldo terbaru
	inventoryBalance, err := repo.InventoryRepository.CurrentBalance(inventoryIn.Code)
	if err != nil {
		return err
	}
	// kalkulasi untuk penentuan harga dan saldo terbaru

	qtyBalance := inventoryBalance.Quantity + inventoryIn.Qty
	amountBalance := inventoryBalance.Amount + (inventoryIn.Price * float64(inventoryIn.Qty))
	priceBalance := amountBalance / float64(qtyBalance)

	// masukkan data terbaru
	data := entities.InventoryIn{
		Code:          inventoryIn.Code,
		Qty:           inventoryIn.Qty,
		Price:         inventoryIn.Price,
		Amount:        inventoryIn.Price * float64(inventoryIn.Qty),
		BalanceQty:    qtyBalance,
		BalancePrice:  priceBalance,
		BalanceAmount: amountBalance,
	}

	if err := repo.InventoryRepository.In(data); err != nil {
		return err
	}
	return nil
}

func (repo repositoryInventoryService) DecreaseStock(inventoryOut *dto.DecreaseInventoryReq) error {
	// cari saldo sekarang
	inventoryBalance, err := repo.InventoryRepository.CurrentBalance(inventoryOut.Code)
	if err != nil {
		return err
	}

	// kondisional pembelian
	if inventoryOut.Qty > inventoryBalance.Quantity {
		return errors.New("persediaan kurang")
	}

	// kalkulasi saldo terbaru
	balanceQty := inventoryBalance.Quantity - inventoryOut.Qty
	balancePrice := inventoryBalance.Price
	balanceAmount := inventoryBalance.Amount - (inventoryBalance.Price * float64(inventoryOut.Qty))

	// kondisional meng nol kan harga
	if balanceQty == 0 {
		balancePrice = 0
		balanceAmount = 0
	}

	// masukkan data ke db
	data := entities.InventoryOut{
		Code:          inventoryOut.Code,
		Price:         inventoryBalance.Price,
		Qty:           inventoryOut.Qty,
		Amount:        inventoryBalance.Price * float64(inventoryOut.Qty),
		BalanceQty:    balanceQty,
		BalancePrice:  balancePrice,
		BalanceAmount: balanceAmount,
	}

	if err := repo.InventoryRepository.Out(data); err != nil {
		return err
	}
	return nil
}
