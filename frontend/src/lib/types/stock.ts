export interface StockData {
	date: Date;
	open: number;
	high: number;
	low: number;
	close: number;
	volume?: number;
}

export interface StockResponse {
	symbol: string;
	data: StockData[];
	metadata: {
		currency: string;
		interval: string;
		lastRefreshed: Date;
		timeZone: string;
	};
	overview?: CompanyOverview;
}

export interface CompanyOverview {
	symbol?: string;
	name?: string;
	exchange?: string;
	sector?: string;
	industry?: string;
	description?: string;
	country?: string;
	currency?: string;
	market_cap?: number;
	pe_ratio?: number;
	forward_pe_ratio?: number;
	eps?: number;
	dividend_per_share?: number;
	dividend_yield?: number;
	beta?: number;
	revenue_ttm?: number;
	profit_margin?: number;
	roa_ttm?: number;
	roe_ttm?: number;
	analyst_target_price?: number;
	week_52_high?: number;
	week_52_low?: number;
	book_value?: number;
	price_to_book?: number;
	shares_outstanding?: number;
}

export interface ChartDimensions {
	width: number;
	height: number;
	margin: {
		top: number;
		right: number;
		bottom: number;
		left: number;
	};
}

export interface TooltipData {
	date: Date;
	open: number;
	high: number;
	low: number;
	close: number;
	volume?: number;
	x: number;
	y: number;
}
