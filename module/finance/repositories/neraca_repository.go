package repositories

import (
	"fmt"

	"github.com/riyan-eng/api-finance/module/finance/repositories/models"
	"github.com/riyan-eng/api-finance/module/finance/services/entities"
	"gorm.io/gorm"
)

type NeracaRepository interface {
	ChartOfAccountParent() ([]models.ChartOfAccountParentModel, error)
	ChartOfAccountChild() ([]models.ChartOfAccountChildModel, error)
	TrialBalance(nE entities.NeracaEntity) ([]entities.NeracaEntity, error)
	TrialBalanceAfterAdjustment(nE entities.NeracaEntity) ([]entities.NeracaEntity, error)
}

type databaseNeracaRepository struct {
	DB *gorm.DB
}

func NewNeracaRepository(DB *gorm.DB) NeracaRepository {
	return &databaseNeracaRepository{
		DB: DB,
	}
}

func (database *databaseNeracaRepository) ChartOfAccountParent() ([]models.ChartOfAccountParentModel, error) {
	chartOfAccountParent := new([]models.ChartOfAccountParentModel)
	query := `
		select c.code, c.name, c.name_bahasa from finance.coas c where c.parent is null order by c.code asc
	`

	if err := database.DB.Raw(query).Scan(&chartOfAccountParent).Error; err != nil {
		return *chartOfAccountParent, err
	}
	return *chartOfAccountParent, nil
}

func (database *databaseNeracaRepository) ChartOfAccountChild() ([]models.ChartOfAccountChildModel, error) {
	chartOfAccountChild := new([]models.ChartOfAccountChildModel)
	query := `
		select c.parent, c.code, c.name, c.name_bahasa from finance.coas c where c.parent is not null order by c.code asc   
	`
	if err := database.DB.Raw(query).Scan(&chartOfAccountChild).Error; err != nil {
		return *chartOfAccountChild, err
	}
	return *chartOfAccountChild, nil
}

func (database *databaseNeracaRepository) TrialBalance(nE entities.NeracaEntity) ([]entities.NeracaEntity, error) {
	neraca := new([]models.NeracaModel)
	query := fmt.Sprint(`
	select code, coa, coalesce(sum(debet), 0) as debet, coalesce(sum(credit), 0) as credit, (coalesce(sum(debet), 0)-coalesce(sum(credit), 0)) as saldo from (
		select c.code as code, c."name" as coa, crj.debet as debet, crj.credit as credit from finance.cash_receipt_journals crj
		join finance.coas c on c.code = crj.coa 
		union all
		select c.code as code, c."name" coa, cpj.debet as debet, cpj.credit as credit from finance.cash_payment_journals cpj
		join finance.coas c on c.code = cpj.coa
		union all
		select c.code as code, c."name" coa, sj.debet as debet, sj.credit as credit from finance.sales_journals sj
		join finance.coas c on c.code = sj.coa
		union all
		select c.code as code, c."name" coa, pj.debet as debet, pj.credit as credit from finance.purchase_journals pj
		join finance.coas c on c.code = pj.coa
		union all
		select c.code as code, c."name" coa, gl.debet as debet, gl.credit as credit from finance.general_ledgers gl
		join finance.coas c on c.code = gl.coa
	) x group by code, coa order by code asc
	`)
	if err := database.DB.Raw(query).Scan(&neraca).Error; err != nil {
		return nil, err
	}

	var neracasEntity []entities.NeracaEntity
	for _, val := range *neraca {
		neracaEntity := entities.NeracaEntity{
			Name:   val.Coa,
			Code:   val.Code,
			Debet:  val.Debet,
			Credit: val.Credit,
			Saldo:  val.Saldo,
		}

		neracasEntity = append(neracasEntity, neracaEntity)
	}

	return neracasEntity, nil
}

func (database *databaseNeracaRepository) TrialBalanceAfterAdjustment(nE entities.NeracaEntity) ([]entities.NeracaEntity, error) {
	neraca := new([]models.NeracaModel)
	query := fmt.Sprint(`
	select code, coa, coalesce(sum(debet), 0) as debet, coalesce(sum(credit), 0) as credit, (coalesce(sum(debet), 0)-coalesce(sum(credit), 0)) as saldo from (
		select c.code as code, c."name" as coa, crj.debet as debet, crj.credit as credit from finance.cash_receipt_journals crj
		join finance.coas c on c.code = crj.coa 
		union all
		select c.code as code, c."name" coa, cpj.debet as debet, cpj.credit as credit from finance.cash_payment_journals cpj
		join finance.coas c on c.code = cpj.coa
		union all
		select c.code as code, c."name" coa, sj.debet as debet, sj.credit as credit from finance.sales_journals sj
		join finance.coas c on c.code = sj.coa
		union all
		select c.code as code, c."name" coa, pj.debet as debet, pj.credit as credit from finance.purchase_journals pj
		join finance.coas c on c.code = pj.coa
		union all
		select c.code as code, c."name" coa, gl.debet as debet, gl.credit as credit from finance.general_ledgers gl
		join finance.coas c on c.code = gl.coa
		union all 
		select c.code as code, c."name" coa, aje.debet as debet, aje.credit as credit from finance.adjustment_journal_entries aje
		join finance.coas c on c.code = aje.coa 
	) x group by code, coa order by code asc
	`)
	if err := database.DB.Raw(query).Scan(&neraca).Error; err != nil {
		return nil, err
	}

	var neracasEntity []entities.NeracaEntity
	for _, val := range *neraca {
		neracaEntity := entities.NeracaEntity{
			Name:   val.Coa,
			Code:   val.Code,
			Debet:  val.Debet,
			Credit: val.Credit,
			Saldo:  val.Saldo,
		}

		neracasEntity = append(neracasEntity, neracaEntity)
	}

	return neracasEntity, nil
}
