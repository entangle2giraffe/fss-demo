package entities

type CompanyOverview struct {
	Symbol             string  `json:"symbol"`
	Name               string  `json:"name"`
	Exchange           string  `json:"exchange"`
	Sector             string  `json:"sector"`
	Industry           string  `json:"industry"`
	Description        string  `json:"description"`
	Country            string  `json:"country"`
	Currency           string  `json:"currency"`
	MarketCap          float64 `json:"market_cap"`
	PERatio            float64 `json:"pe_ratio"`
	ForwardPERatio     float64 `json:"forward_pe_ratio"`
	EPS                float64 `json:"eps"`
	DividendPerShare   float64 `json:"dividend_per_share"`
	DividendYield      float64 `json:"dividend_yield"`
	Beta               float64 `json:"beta"`
	RevenueTTM         float64 `json:"revenue_ttm"`
	ProfitMargin       float64 `json:"profit_margin"`
	ReturnOnAssetsTTM  float64 `json:"roa_ttm"`
	ReturnOnEquityTTM  float64 `json:"roe_ttm"`
	AnalystTargetPrice float64 `json:"analyst_target_price"`
	Week52High         float64 `json:"week_52_high"`
	Week52Low          float64 `json:"week_52_low"`
	BookValue          float64 `json:"book_value"`
	PriceToBookRatio   float64 `json:"price_to_book"`
	SharesOutstanding  float64 `json:"shares_outstanding"`
}

type StockWithOverview struct {
	Symbol   string               `json:"symbol"`
	Data     map[string]StockData `json:"data"`
	Overview CompanyOverview      `json:"overview"`
}
