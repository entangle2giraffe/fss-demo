package alphavantage

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"stock-ingestor/internal/domain/entities"
	"stock-ingestor/internal/domain/repositories"
)

type StockRepositoryImpl struct {
	adapter *AlphaVantageAdapter
}

func NewStockRepositoryImpl(apiKey, url string) repositories.AlphaVantageRepository {
	return &StockRepositoryImpl{
		adapter: NewAlphaVantageAdapter(apiKey, url),
	}
}

func (sri *StockRepositoryImpl) GetTimeSeries(symbol string, period string) (*entities.StockTimeSeries, error) {
	var p int
	switch period {
	case "daily":
		p = Daily
	case "weekly":
		p = Weekly
	case "monthly", "yearly":
		p = Monthly
	default:
		p = Weekly
	}

	// Use non-adjusted endpoints to avoid premium requirement
	resp, err := sri.adapter.GetTimeSeries(p, false, symbol, false)
	if err != nil {
		return nil, fmt.Errorf("failed to get time series: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	ts, err := sri.adapter.ParseTimeSeriesResponse(resp)
	if err != nil {
		return nil, err
	}

	// For yearly, aggregate the monthly data into yearly OHLCV
	if period == "yearly" {
		ts = aggregateToYearly(ts)
	}

	return ts, nil
}

func (sri *StockRepositoryImpl) GetOverview(symbol string) (entities.CompanyOverview, error) {
	overviewPayload, err := sri.adapter.GetOverview(symbol)
	if err != nil {
		return entities.CompanyOverview{}, err
	}

	return toCompanyOverview(overviewPayload), nil
}

// aggregateToYearly collapses monthly data into yearly OHLCV keyed by Dec 31 of each year.
func aggregateToYearly(ts *entities.StockTimeSeries) *entities.StockTimeSeries {
	byYear := make(map[int][]entities.StockData)

	for _, v := range ts.Data {
		y := v.Date.Year()
		byYear[y] = append(byYear[y], v)
	}

	aggregated := make(map[string]entities.StockData)
	for year, records := range byYear {
		if len(records) == 0 {
			continue
		}

		// Sort by date to pick open/close correctly
		sorted := records
		for i := 0; i < len(sorted)-1; i++ {
			for j := i + 1; j < len(sorted); j++ {
				if sorted[j].Date.Before(sorted[i].Date) {
					sorted[i], sorted[j] = sorted[j], sorted[i]
				}
			}
		}

		open := sorted[0].Open
		close := sorted[len(sorted)-1].Close
		high := sorted[0].High
		low := sorted[0].Low
		var volume int64

		for _, r := range sorted {
			if r.High > high {
				high = r.High
			}
			if r.Low < low {
				low = r.Low
			}
			volume += r.Volume
		}

		keyDate := time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC)
		aggregated[keyDate.Format("2006-01-02")] = entities.StockData{
			Date:   keyDate,
			Open:   open,
			High:   high,
			Low:    low,
			Close:  close,
			Volume: volume,
		}
	}

	return &entities.StockTimeSeries{Symbol: ts.Symbol, Data: aggregated}
}

// toCompanyOverview maps Alpha Vantage overview payload into typed struct with best-effort parsing.
func toCompanyOverview(m map[string]interface{}) entities.CompanyOverview {
	getS := func(key string) string {
		if v, ok := m[key].(string); ok {
			return v
		}
		return ""
	}
	getF := func(key string) float64 {
		s := getS(key)
		if s == "" {
			return 0
		}
		if f, err := strconv.ParseFloat(s, 64); err == nil {
			return f
		}
		return 0
	}

	return entities.CompanyOverview{
		Symbol:             getS("Symbol"),
		Name:               getS("Name"),
		Exchange:           getS("Exchange"),
		Sector:             getS("Sector"),
		Industry:           getS("Industry"),
		Description:        getS("Description"),
		Country:            getS("Country"),
		Currency:           getS("Currency"),
		MarketCap:          getF("MarketCapitalization"),
		PERatio:            getF("PERatio"),
		ForwardPERatio:     getF("ForwardPE"),
		EPS:                getF("EPS"),
		DividendPerShare:   getF("DividendPerShare"),
		DividendYield:      getF("DividendYield"),
		Beta:               getF("Beta"),
		RevenueTTM:         getF("RevenueTTM"),
		ProfitMargin:       getF("ProfitMargin"),
		ReturnOnAssetsTTM:  getF("ReturnOnAssetsTTM"),
		ReturnOnEquityTTM:  getF("ReturnOnEquityTTM"),
		AnalystTargetPrice: getF("AnalystTargetPrice"),
		Week52High:         getF("52WeekHigh"),
		Week52Low:          getF("52WeekLow"),
		BookValue:          getF("BookValue"),
		PriceToBookRatio:   getF("PriceToBookRatio"),
		SharesOutstanding:  getF("SharesOutstanding"),
	}
}
