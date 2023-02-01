package repositories

import (
	"fmt"

	"github.com/riyan-eng/api-finance/module/finance/repositories/models"
	"github.com/riyan-eng/api-finance/module/finance/services/entities"
	"gorm.io/gorm"
)

type InventoryRepository interface {
	CurrentBalance(code string) (*models.InventoryBalance, error)
	In(entities.InventoryIn) error
	Out(entities.InventoryOut) error
}

type databaseInventoryRepository struct {
	DB *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) InventoryRepository {
	return &databaseInventoryRepository{
		DB: db,
	}
}

func (database *databaseInventoryRepository) CurrentBalance(code string) (*models.InventoryBalance, error) {
	data := new(models.InventoryBalance)
	query := fmt.Sprintf(`
	select gs.saldo_quantity as quantity, gs.saldo_price as price, gs.saldo_amount as amount from finance.good_stocks gs where gs.good='%s' order by gs.date_time desc limit 1
	`, code)
	if err := database.DB.Raw(query).Scan(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (database *databaseInventoryRepository) In(inventoryIn entities.InventoryIn) error {
	query := fmt.Sprintf(`
	insert into finance.good_stocks(good, dk, quantity, price, amount, saldo_quantity, saldo_price, saldo_amount)
	values ('%v', 'D', %v, %f, %f, %v, %f, %f )
	`, inventoryIn.Code, inventoryIn.Qty, inventoryIn.Price, inventoryIn.Amount, inventoryIn.BalanceQty, inventoryIn.BalancePrice, inventoryIn.BalanceAmount)

	if err := database.DB.Exec(query).Error; err != nil {
		return err
	}

	return nil
}

func (database *databaseInventoryRepository) Out(inventoryOut entities.InventoryOut) error {
	query := fmt.Sprintf(`
	insert into finance.good_stocks(good, dk, quantity, price, amount, saldo_quantity, saldo_price, saldo_amount)
	values ('%v', 'K', %v, %f, %f, %v, %f, %f )
	`, inventoryOut.Code, inventoryOut.Qty, inventoryOut.Price, inventoryOut.Amount, inventoryOut.BalanceQty, inventoryOut.BalancePrice, inventoryOut.BalanceAmount)

	if err := database.DB.Exec(query).Error; err != nil {
		return err
	}

	return nil
}
