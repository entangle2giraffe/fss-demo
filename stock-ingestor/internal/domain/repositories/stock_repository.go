package repositories

import (
	"stock-ingestor/internal/domain/entities"
)

type StockRepository interface {
	GetTimeSeries(symbol string, period string) (*entities.StockTimeSeries, error)
	GetOverview(symbol string) (entities.CompanyOverview, error)
}
