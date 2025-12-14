<script lang="ts">
	import formatNumber from '$lib/utils';

	interface Props {
		companyData: {
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
		volume?: number;
		avgVolume?: number;
	}

	let { companyData, volume, avgVolume }: Props = $props();
</script>

<div class="bg-white dark:bg-gray-800 rounded-lg shadow-lg p-6">
	<div class="grid grid-cols-1 md:grid-cols-2 gap-6">
		<div class="space-y-3">
			<div
				class="flex justify-between items-center py-2 border-b border-gray-100 dark:border-gray-700"
			>
				<span class="text-sm font-medium text-gray-600 dark:text-gray-400">Market Cap</span>
				<span class="font-semibold text-gray-900 dark:text-gray-100"
					>{formatNumber(companyData.marketCap)}</span
				>
			</div>
			<div
				class="flex justify-between items-center py-2 border-b border-gray-100 dark:border-gray-700"
			>
				<span class="text-sm font-medium text-gray-600 dark:text-gray-400">P/E Ratio</span>
				<span class="font-semibold text-gray-900 dark:text-gray-100">{companyData.peRatio}</span>
			</div>
			<div
				class="flex justify-between items-center py-2 border-b border-gray-100 dark:border-gray-700"
			>
				<span class="text-sm font-medium text-gray-600 dark:text-gray-400">EPS</span>
				<span class="font-semibold text-gray-900 dark:text-gray-100">${companyData.eps}</span>
			</div>
			<div
				class="flex justify-between items-center py-2 border-b border-gray-100 dark:border-gray-700"
			>
				<span class="text-sm font-medium text-gray-600 dark:text-gray-400">Dividend Yield</span>
				<span class="font-semibold text-gray-900 dark:text-gray-100"
					>{(companyData.dividendYield * 100).toFixed(2)}%</span
				>
			</div>
			<div class="flex justify-between items-center py-2">
				<span class="text-sm font-medium text-gray-600 dark:text-gray-400">Exchange</span>
				<span class="font-semibold text-gray-900 dark:text-gray-100">{companyData.exchange}</span>
			</div>
		</div>
		<div class="space-y-3">
			<div
				class="flex justify-between items-center py-2 border-b border-gray-100 dark:border-gray-700"
			>
				<span class="text-sm font-medium text-gray-600 dark:text-gray-400">52W Range</span>
				<span class="font-semibold text-gray-900 dark:text-gray-100"
					>${companyData.week52Low.toFixed(2)} - ${companyData.week52High.toFixed(2)}</span
				>
			</div>
			<div
				class="flex justify-between items-center py-2 border-b border-gray-100 dark:border-gray-700"
			>
				<span class="text-sm font-medium text-gray-600 dark:text-gray-400">Volume</span>
				<span class="font-semibold text-gray-900 dark:text-gray-100"
					>{formatNumber(volume ?? 0)}</span
				>
			</div>
			<div
				class="flex justify-between items-center py-2 border-b border-gray-100 dark:border-gray-700"
			>
				<span class="text-sm font-medium text-gray-600 dark:text-gray-400">Avg Volume</span>
				<span class="font-semibold text-gray-900 dark:text-gray-100"
					>{formatNumber(avgVolume ?? 0)}</span
				>
			</div>
			<div
				class="flex justify-between items-center py-2 border-b border-gray-100 dark:border-gray-700"
			>
				<span class="text-sm font-medium text-gray-600 dark:text-gray-400">Beta</span>
				<span class="font-semibold text-gray-900 dark:text-gray-100">{companyData.beta}</span>
			</div>
			<div class="flex justify-between items-center py-2">
				<span class="text-sm font-medium text-gray-600 dark:text-gray-400">Sector</span>
				<span class="font-semibold text-gray-900 dark:text-gray-100">{companyData.sector}</span>
			</div>
		</div>
	</div>
</div>
