package usecases

import (
	"stock-ingestor/internal/domain/entities"
	"stock-ingestor/internal/domain/repositories"
)

type GetCompanyOverviewUseCase struct {
	repo repositories.StockRepository
}

func NewGetCompanyOverviewUseCase(repo repositories.StockRepository) *GetCompanyOverviewUseCase {
	return &GetCompanyOverviewUseCase{repo: repo}
}

func (uc *GetCompanyOverviewUseCase) Execute(symbol string) (entities.CompanyOverview, error) {
	return uc.repo.GetOverview(symbol)
}
