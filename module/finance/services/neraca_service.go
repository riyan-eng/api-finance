package services

import (
	"github.com/riyan-eng/api-finance/module/finance/repositories"
	"github.com/riyan-eng/api-finance/module/finance/services/entities"
)

type NeracaService interface {
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
