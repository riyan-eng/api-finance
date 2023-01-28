package services

import (
	"github.com/riyan-eng/api-finance/module/finance/repositories"
	"github.com/riyan-eng/api-finance/module/finance/services/entities"
)

type NeracaService interface {
	ChartOfAccount() ([]entities.COA, error)
	TrialBalance(nE entities.NeracaEntity) ([]entities.NeracaEntity, error)
	TrialBalanceAfterAdjustment(nE entities.NeracaEntity) ([]entities.NeracaEntity, error)
}

type neracaRepository struct {
	NeracaRepository repositories.NeracaRepository
}

func NewNeracaService(nR repositories.NeracaRepository) NeracaService {
	return &neracaRepository{
		NeracaRepository: nR,
	}
}

func (repo *neracaRepository) ChartOfAccount() ([]entities.COA, error) {
	var data []entities.COA
	coaParents, err := repo.NeracaRepository.ChartOfAccountParent()
	if err != nil {
		return data, nil
	}
	// fmt.Println(coaParent)
	coaChilds, err := repo.NeracaRepository.ChartOfAccountChild()
	if err != nil {
		return data, nil
	}
	// fmt.Println(coaChild)

	for _, parent := range coaParents {
		coaParent := entities.COA{
			Code: parent.Code,
			Name: parent.Name,
		}
		// fmt.Println(parent)
		var newCoaChilds []entities.COAChild
		for _, child := range coaChilds {
			if parent.Code == child.Parent {
				// fmt.Println(child)
				coaChild := entities.COAChild{
					Code:       child.Code,
					Name:       child.Name,
					NameBahasa: child.NameBahasa,
				}
				newCoaChilds = append(newCoaChilds, coaChild)

			}
		}
		coaParent.Child = newCoaChilds
		data = append(data, coaParent)
	}
	// fmt.Println(data)
	return data, nil
}

func (repo *neracaRepository) TrialBalance(nE entities.NeracaEntity) ([]entities.NeracaEntity, error) {
	neracas, err := repo.NeracaRepository.TrialBalance(nE)
	if err != nil {
		return nil, err
	}
	return neracas, nil
}

func (repo *neracaRepository) TrialBalanceAfterAdjustment(nE entities.NeracaEntity) ([]entities.NeracaEntity, error) {
	neracas, err := repo.NeracaRepository.TrialBalanceAfterAdjustment(nE)
	if err != nil {
		return nil, err
	}
	return neracas, nil
}
