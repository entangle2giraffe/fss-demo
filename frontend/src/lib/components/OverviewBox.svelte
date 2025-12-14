<script lang="ts">
	import type { StockResponse } from '$lib/types/stock';
	import StockHeader from './StockHeader.svelte';
	import TabNavigation from './TabNavigation.svelte';
	import OverviewContent from './OverviewContent.svelte';
	import FinancialsContent from './FinancialsContent.svelte';
	import AnalysisContent from './AnalysisContent.svelte';

	import '$lib/components/overview-style.css';

	interface Props {
		stockData?: StockResponse;
	}

	let { stockData }: Props = $props();

	// Tab state
	let activeTab = $state('overview');

	// Calculate metrics from stock data
	const currentPrice = $derived(stockData?.data?.[stockData.data.length - 1]?.close || 0);
	const previousPrice = $derived(
		stockData?.data?.[stockData.data.length - 2]?.close || currentPrice
	);
	const priceChange = $derived(currentPrice - previousPrice);
	const priceChangePercent = $derived(previousPrice > 0 ? (priceChange / previousPrice) * 100 : 0);
	const volume = $derived(stockData?.data?.[stockData.data.length - 1]?.volume || 0);
	const avgVolume = $derived(
		stockData?.data
			? stockData.data.reduce((sum, d) => sum + (d.volume || 0), 0) / stockData.data.length
			: 0
	);

	type CompanyData = {
		symbol: string;
		name: string;
		exchange: string;
		sector: string;
		marketCap: number;
		peRatio: number;
		forwardPE: number;
		eps: number;
		dividendYield: number;
		beta: number;
		revenueTTM: number;
		profitMargin: number;
		roa: number;
		roe: number;
		analystTarget: number;
		week52High: number;
		week52Low: number;
		dividendPerShare: number;
		bookValue: number;
		priceToBook: number;
		sharesOutstanding: number;
	};

	const fallbackNames: Record<string, string> = {
		AAPL: 'Apple Inc.',
		GOOGL: 'Alphabet Inc.',
		MSFT: 'Microsoft Corporation',
		TSLA: 'Tesla, Inc.',
		AMZN: 'Amazon.com, Inc.',
		META: 'Meta Platforms, Inc.',
		NVDA: 'NVIDIA Corporation',
		JPM: 'JPMorgan Chase & Co.',
		V: 'Visa Inc.',
		WMT: 'Walmart Inc.'
	};

	const companyData: CompanyData = $derived((() => {
		const baseSymbol = stockData?.symbol || 'N/A';
		const base: CompanyData = {
			symbol: baseSymbol,
			name:
				(baseSymbol && fallbackNames[baseSymbol])
					? fallbackNames[baseSymbol]
					: baseSymbol !== 'N/A'
						? `${baseSymbol} Inc.`
						: 'Company Name',
			exchange: 'N/A',
			sector: 'N/A',
			marketCap: 0,
			peRatio: 0,
			forwardPE: 0,
			eps: 0,
			dividendYield: 0,
			beta: 0,
			revenueTTM: 0,
			profitMargin: 0,
			roa: 0,
			roe: 0,
			analystTarget: 0,
			week52High: 0,
			week52Low: 0,
			dividendPerShare: 0,
			bookValue: 0,
			priceToBook: 0,
			sharesOutstanding: 0
		};

		const overview = stockData?.overview;
		if (!overview) return base;

		return {
			...base,
			symbol: overview.symbol || base.symbol,
			name: overview.name || base.name,
			exchange: overview.exchange || base.exchange,
			sector: overview.sector || base.sector,
			marketCap: overview.market_cap ?? base.marketCap,
			peRatio: overview.pe_ratio ?? base.peRatio,
			forwardPE: overview.forward_pe_ratio ?? base.forwardPE,
			eps: overview.eps ?? base.eps,
			dividendYield: overview.dividend_yield ?? base.dividendYield,
			beta: overview.beta ?? base.beta,
			revenueTTM: overview.revenue_ttm ?? base.revenueTTM,
			profitMargin: overview.profit_margin ?? base.profitMargin,
			roa: overview.roa_ttm ?? base.roa,
			roe: overview.roe_ttm ?? base.roe,
			analystTarget: overview.analyst_target_price ?? base.analystTarget,
			week52High: overview.week_52_high ?? base.week52High,
			week52Low: overview.week_52_low ?? base.week52Low,
			dividendPerShare: overview.dividend_per_share ?? base.dividendPerShare,
			bookValue: overview.book_value ?? base.bookValue,
			priceToBook: overview.price_to_book ?? base.priceToBook,
			sharesOutstanding: overview.shares_outstanding ?? base.sharesOutstanding
		};
	})());

	function switchTab(tab: string) {
		activeTab = tab;
	}

	// Get current tab component
	const currentTabComponent = $derived(
		activeTab === 'overview'
			? OverviewContent
			: activeTab === 'financials'
				? FinancialsContent
				: AnalysisContent
	);
</script>

<div class="company-overview space-y-6">
	<!-- Section 1: Stock Header (Always Visible) -->
	<section class="stock-header-section">
		<StockHeader {companyData} {currentPrice} {priceChange} {priceChangePercent} />
	</section>

	<!-- Section 2: Tab Navigation -->
	<section class="tab-navigation-section">
		<TabNavigation {activeTab} onTabChange={switchTab} />
	</section>

	<!-- Section 3: Tab Content (Dynamic) -->
	<section class="tab-content-section">
		{#if activeTab === 'overview'}
			<OverviewContent {companyData} {volume} {avgVolume} />
		{:else if activeTab === 'financials'}
			<FinancialsContent {companyData} />
		{:else if activeTab === 'analysis'}
			<AnalysisContent {companyData} {currentPrice} />
		{/if}
	</section>
</div>
