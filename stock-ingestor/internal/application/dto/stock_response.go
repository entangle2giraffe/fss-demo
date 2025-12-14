package dto

import (
	"stock-ingestor/internal/domain/entities"
)

type StockTimeSeriesResponse struct {
	Symbol string                        `json:"symbol"`
	Data   map[string]entities.StockData `json:"data"`
}
