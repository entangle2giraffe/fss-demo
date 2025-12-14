package repositories

import (
	"stock-ingestor/internal/domain/entities"
)

// StockReadRepository for reading stock data from database
type StockReadRepository interface {
	GetTimeSeries(symbol string, period string) (*entities.StockTimeSeries, error)
	GetOverview(symbol string) (entities.CompanyOverview, error)
}

// StockWriteRepository for writing stock data to database
type StockWriteRepository interface {
	StoreTimeSeries(symbol string, data *entities.StockTimeSeries) error
	StoreOverview(symbol string, overview entities.CompanyOverview) error
}

// AlphaVantageRepository for fetching data from AlphaVantage API
type AlphaVantageRepository interface {
	GetTimeSeries(symbol string, period string) (*entities.StockTimeSeries, error)
	GetOverview(symbol string) (entities.CompanyOverview, error)
}
