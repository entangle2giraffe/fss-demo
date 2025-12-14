package entities

import "time"

type StockData struct {
	Date           time.Time `json:"date"`
	Open           float64   `json:"open"`
	High           float64   `json:"high"`
	Low            float64   `json:"low"`
	Close          float64   `json:"close"`
	AdjustedClose  float64   `json:"adjusted_close"`
	Volume         int64     `json:"volume"`
	DividendAmount float64   `json:"dividend_amount"`
}

type StockTimeSeries struct {
	Symbol string               `json:"symbol"`
	Data   map[string]StockData `json:"data"`
}
