package usecases

import (
	"stock-ingestor/internal/domain/entities"
	"stock-ingestor/internal/domain/repositories"
)

type GetStockTimeSeriesUseCase struct {
	repo repositories.StockRepository
}

func NewGetStockTimeSeriesUseCase(repo repositories.StockRepository) *GetStockTimeSeriesUseCase {
	return &GetStockTimeSeriesUseCase{
		repo: repo,
	}
}

func (uc *GetStockTimeSeriesUseCase) Execute(symbol, period string) (*entities.StockTimeSeries, error) {
	return uc.repo.GetTimeSeries(symbol, period)
}
