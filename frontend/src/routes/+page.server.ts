import type { StockResponse, StockData } from '$lib/types/stock';

export async function load({ url, fetch }: { url: URL, fetch: any }) {
	const symbol = (url.searchParams.get('symbol') || 'AAPL').toUpperCase();
	const period = url.searchParams.get('period') || 'daily';

	try {
		const timeSeriesUrl = `/api/stocks/${symbol}?period=${period}`;
		const overviewUrl = `/api/stocks/${symbol}/overview`;

		const [timeSeriesResp, overviewResp] = await Promise.all([
			fetch(timeSeriesUrl),
			fetch(overviewUrl)
		]);

		if (!timeSeriesResp.ok) {
			console.error('Failed to fetch stock data:', timeSeriesResp.status, timeSeriesResp.statusText);
			// Return empty data on error
			return {
				stockData: {
					symbol,
					data: [],
					metadata: {
						currency: 'USD',
						interval: period,
						lastRefreshed: new Date(),
						timeZone: 'America/New_York'
					}
				},
				symbols: ['AAPL', 'GOOGL', 'MSFT', 'TSLA', 'AMZN', 'META', 'NVDA'],
				selectedSymbol: symbol,
				selectedPeriod: period
			};
		}

		const body = await timeSeriesResp.json();
		const stockDataArray: StockData[] = body.data ?? body;

		const overview = overviewResp.ok ? await overviewResp.json() : null;

		// Ensure dates are Date objects and sorted ascending for the charts
		const normalized = (stockDataArray || [])
			.map((d) => ({
				...d,
				date: new Date(d.date)
			}))
			.sort((a, b) => a.date.getTime() - b.date.getTime());

		const stockResponse: StockResponse = {
			symbol,
			data: normalized,
			overview: overview ?? undefined,
			metadata: {
				currency: 'USD',
				interval: period,
				lastRefreshed: new Date(),
				timeZone: 'America/New_York'
			}
		};

		return {
			stockData: stockResponse,
			symbols: ['AAPL', 'GOOGL', 'MSFT', 'TSLA', 'AMZN', 'META', 'NVDA'],
			selectedSymbol: symbol,
			selectedPeriod: period
		};
	} catch (error) {
		console.error('Error loading stock data:', error);
		
		// Return empty data on error
		return {
			stockData: {
				symbol,
				data: [],
				metadata: {
					currency: 'USD',
					interval: period,
					lastRefreshed: new Date(),
					timeZone: 'America/New_York'
				}
			},
			symbols: ['AAPL', 'GOOGL', 'MSFT', 'TSLA', 'AMZN', 'META', 'NVDA'],
			selectedSymbol: symbol,
			selectedPeriod: period
		};
	}
}
