package repositories

import (
	"github.com/riyan-eng/api-finance/database"
	"github.com/riyan-eng/api-finance/src/finance/repositories/models"
)

// type CashRepositoryInterface struct{
// 	InsertOne()
// 	Sum()
// }

var db = database.DB

func Save(product models.CashModel) models.CashModel {
	db.Exec(`
	INSERT 
	`)
	return product
}
