package alphavantage

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"stock-ingestor/internal/domain/entities"
)

const (
	Daily = iota
	Weekly
	Monthly
)

type AlphaVantageAdapter struct {
	apiKey string
	url    string
	client *http.Client
}

func NewAlphaVantageAdapter(apiKey string, url string) *AlphaVantageAdapter {
	return &AlphaVantageAdapter{
		apiKey: apiKey,
		url:    url,
		client: &http.Client{Timeout: 15 * time.Second},
	}
}

func (aaa *AlphaVantageAdapter) GetTimeSeries(period int, fullOutput bool, symbol string, adjusted bool) (*http.Response, error) {
	var apiFn string
	switch period {
	case Daily:
		apiFn = "TIME_SERIES_DAILY"
	case Weekly:
		apiFn = "TIME_SERIES_WEEKLY"
	case Monthly:
		apiFn = "TIME_SERIES_MONTHLY"
	}

	if adjusted {
		apiFn += "_ADJUSTED"
	}

	apiUrl := aaa.url + fmt.Sprintf("function=%s&symbol=%s&apikey=%s", apiFn, symbol, aaa.apiKey)

	if fullOutput && period == Daily {
		apiUrl += "&outputsize=full"
	}

	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	return resp, nil
}

// GetOverview calls the Alpha Vantage OVERVIEW endpoint
func (aaa *AlphaVantageAdapter) GetOverview(symbol string) (map[string]interface{}, error) {
	apiUrl := fmt.Sprintf("%sfunction=OVERVIEW&symbol=%s&apikey=%s", aaa.url, symbol, aaa.apiKey)
	resp, err := aaa.client.Get(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to make overview request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("overview request failed with status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read overview body: %w", err)
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(body, &payload); err != nil {
		return nil, fmt.Errorf("failed to decode overview: %w; body=%s", err, string(body))
	}

	// Surface errors from AV
	if note, ok := payload["Note"].(string); ok && note != "" {
		return nil, fmt.Errorf("alphavantage overview rate limit: %s", note)
	}
	if emsg, ok := payload["Error Message"].(string); ok && emsg != "" {
		return nil, fmt.Errorf("alphavantage overview error: %s", emsg)
	}
	if info, ok := payload["Information"].(string); ok && info != "" {
		return nil, fmt.Errorf("alphavantage overview information: %s", info)
	}

	log.Printf("alphavantage overview keys=%v", keysOf(payload))

	return payload, nil
}

func (aaa *AlphaVantageAdapter) ParseTimeSeriesResponse(resp *http.Response) (*entities.StockTimeSeries, error) {
	defer resp.Body.Close()

	var rawResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&rawResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Debug keys to surface what AV sent (helps with 500s)
	log.Printf("alphavantage: response keys=%v", keysOf(rawResponse))

	// Alpha Vantage returns rate-limit or error messages under these keys
	if note, ok := rawResponse["Note"].(string); ok && note != "" {
		log.Printf("alphavantage note: %s", note)
		return nil, fmt.Errorf("alphavantage note: %s", note)
	}
	if emsg, ok := rawResponse["Error Message"].(string); ok && emsg != "" {
		log.Printf("alphavantage error: %s", emsg)
		return nil, fmt.Errorf("alphavantage error: %s", emsg)
	}
	if info, ok := rawResponse["Information"].(string); ok && info != "" {
		log.Printf("alphavantage information: %s", info)
		return nil, fmt.Errorf("alphavantage information: %s", info)
	}

	// Find the time series key (works for Daily/Weekly/Monthly, adjusted or not)
	var timeSeriesKey string
	for key := range rawResponse {
		if strings.Contains(strings.ToLower(key), "time series") {
			timeSeriesKey = key
			break
		}
	}

	if timeSeriesKey == "" {
		log.Printf("alphavantage: no time series key found; keys=%v", keysOf(rawResponse))
		return nil, fmt.Errorf("time series data not found in response; keys: %v", keysOf(rawResponse))
	}

	timeSeriesData, ok := rawResponse[timeSeriesKey].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid time series data format")
	}

	data := make(map[string]entities.StockData)
	for dateStr, entry := range timeSeriesData {
		entryMap, ok := entry.(map[string]interface{})
		if !ok {
			continue
		}

		stockData := entities.StockData{}
		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			continue
		}
		stockData.Date = date

		if open, err := parseFloat(entryMap["1. open"]); err == nil {
			stockData.Open = open
		}
		if high, err := parseFloat(entryMap["2. high"]); err == nil {
			stockData.High = high
		}
		if low, err := parseFloat(entryMap["3. low"]); err == nil {
			stockData.Low = low
		}
		if close, err := parseFloat(entryMap["4. close"]); err == nil {
			stockData.Close = close
		}

		// Adjusted close falls back to close if missing (non-adjusted series)
		if adjClose, err := parseFloat(entryMap["5. adjusted close"]); err == nil {
			stockData.AdjustedClose = adjClose
		} else {
			stockData.AdjustedClose = stockData.Close
		}

		// Volume key differs between adjusted/non-adjusted series
		if volume, err := parseInt64(entryMap["6. volume"]); err == nil {
			stockData.Volume = volume
		} else if volume, err := parseInt64(entryMap["5. volume"]); err == nil {
			stockData.Volume = volume
		}

		if dividend, err := parseFloat(entryMap["7. dividend amount"]); err == nil {
			stockData.DividendAmount = dividend
		}

		data[dateStr] = stockData
	}

	symbol := ""
	if meta, ok := rawResponse["Meta Data"].(map[string]interface{}); ok {
		if sym, ok := meta["2. Symbol"].(string); ok {
			symbol = strings.ToUpper(sym)
		}
	}

	return &entities.StockTimeSeries{Symbol: symbol, Data: data}, nil
}

// keysOf returns map keys as []string (for diagnostics)
func keysOf(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func parseFloat(value interface{}) (float64, error) {
	switch v := value.(type) {
	case string:
		return strconv.ParseFloat(v, 64)
	case float64:
		return v, nil
	default:
		return 0, fmt.Errorf("cannot parse float")
	}
}

func parseInt64(value interface{}) (int64, error) {
	switch v := value.(type) {
	case string:
		return strconv.ParseInt(v, 10, 64)
	case float64:
		return int64(v), nil
	default:
		return 0, fmt.Errorf("cannot parse int64")
	}
}
