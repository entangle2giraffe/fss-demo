<script lang="ts">
	import { goto } from '$app/navigation';
	import CandleStickChart from '$lib/components/CandleStickChart.svelte';
	import VolumeBarChart from '$lib/components/VolumeBarChart.svelte';
	import OverviewBox from '$lib/components/OverviewBox.svelte';
	import SymbolSelector from '$lib/components/SymbolSelector.svelte';
	import PeriodSelector from '$lib/components/PeriodSelector.svelte';
	import type { StockResponse } from '$lib/types/stock';

	interface PageData {
		stockData: StockResponse;
		symbols: string[];
		selectedSymbol: string;
		selectedPeriod: string;
	}

	export let data: PageData;

	let selectedSymbol = data.selectedSymbol;
	let selectedPeriod = data.selectedPeriod;
	let stockData = data.stockData;

	// When the load function re-runs (symbol/period change), update local state
	$: {
		selectedSymbol = data.selectedSymbol;
		selectedPeriod = data.selectedPeriod;
		stockData = data.stockData;
	}

	function handleSymbolChange(symbol: string) {
		selectedSymbol = symbol;
		updateUrl();
	}

	function handlePeriodChange(period: string) {
		selectedPeriod = period;
		updateUrl();
	}

	function updateUrl() {
		const url = new URL(window.location.href);
		url.searchParams.set('symbol', selectedSymbol);
		url.searchParams.set('period', selectedPeriod);
		// Trigger navigation so +page.server.ts re-runs and data updates
		void goto(url.toString(), { replaceState: true, noScroll: true });
	}
</script>

<div class="min-h-screen bg-gray-50 dark:bg-gray-900 p-4">
	<div class="max-w-full mx-auto">
		<div class="grid grid-rows-2 lg:grid-cols-2 gap-6 min-h-[calc(100vh-8rem)]">
			<!-- Section 1: Controls + Overview (Top on mobile, Left on desktop) -->
			<div class="space-y-6">
				<!-- Combined Symbol and Period Selector Card -->
				<div class="bg-white dark:bg-gray-800 rounded-lg shadow-lg p-4">
					<div class="space-y-4">
						<SymbolSelector {selectedSymbol} onSymbolChange={handleSymbolChange} />
						<PeriodSelector {selectedPeriod} onPeriodChange={handlePeriodChange} />
					</div>
				</div>

				<!-- Overview Box -->
				<div>
					<OverviewBox {stockData} />
				</div>
			</div>

			<!-- Section 2: Charts (Bottom on mobile, Right on desktop) -->
			<div class="space-y-6">
				<!-- Candlestick Chart -->
				<div>
					<CandleStickChart data={stockData.data} symbol={stockData.symbol} height={400} />
				</div>

				<!-- Volume Bar Chart -->
				<div>
					<VolumeBarChart data={stockData.data} height={120} />
				</div>
			</div>
		</div>
	</div>
</div>
