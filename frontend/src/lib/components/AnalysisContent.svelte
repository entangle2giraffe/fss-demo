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
		currentPrice?: number;
	}

	let { companyData, currentPrice }: Props = $props();

	// Calculate upside potential
	const upsidePotential = $derived(
		currentPrice ? ((companyData.analystTarget - currentPrice) / currentPrice) * 100 : 0
	);
</script>

<div class="bg-white dark:bg-gray-800 rounded-lg shadow-lg p-6">
	<div class="grid grid-cols-1 md:grid-cols-2 gap-6">
		<div class="space-y-3">
			<div
				class="flex justify-between items-center py-2 border-b border-gray-100 dark:border-gray-700"
			>
				<span class="text-sm font-medium text-gray-600 dark:text-gray-400">Analyst Target</span>
				<span class="font-semibold text-gray-900 dark:text-gray-100"
					>${companyData.analystTarget}</span
				>
			</div>
			<div
				class="flex justify-between items-center py-2 border-b border-gray-100 dark:border-gray-700"
			>
				<span class="text-sm font-medium text-gray-600 dark:text-gray-400">Shares Outstanding</span>
				<span class="font-semibold text-gray-900 dark:text-gray-100"
					>{formatNumber(companyData.sharesOutstanding)}</span
				>
			</div>
			<div
				class="flex justify-between items-center py-2 border-b border-gray-100 dark:border-gray-700"
			>
				<span class="text-sm font-medium text-gray-600 dark:text-gray-400">Sector</span>
				<span class="font-semibold text-gray-900 dark:text-gray-100">{companyData.sector}</span>
			</div>
			<div class="flex justify-between items-center py-2">
				<span class="text-sm font-medium text-gray-600 dark:text-gray-400">Upside Potential</span>
				<span class="font-semibold text-green-600 dark:text-green-400"
					>{upsidePotential.toFixed(1)}%</span
				>
			</div>
		</div>
		<div class="space-y-3">
			<div
				class="flex justify-between items-center py-2 border-b border-gray-100 dark:border-gray-700"
			>
				<span class="text-sm font-medium text-gray-600 dark:text-gray-400">50-Day MA</span>
				<span class="font-semibold text-gray-900 dark:text-gray-100">295.96</span>
			</div>
			<div
				class="flex justify-between items-center py-2 border-b border-gray-100 dark:border-gray-700"
			>
				<span class="text-sm font-medium text-gray-600 dark:text-gray-400">200-Day MA</span>
				<span class="font-semibold text-gray-900 dark:text-gray-100">267.22</span>
			</div>
			<div
				class="flex justify-between items-center py-2 border-b border-gray-100 dark:border-gray-700"
			>
				<span class="text-sm font-medium text-gray-600 dark:text-gray-400">Avg Volume</span>
				<span class="font-semibold text-gray-900 dark:text-gray-100">{formatNumber(45000000)}</span>
			</div>
			<div class="flex justify-between items-center py-2">
				<span class="text-sm font-medium text-gray-600 dark:text-gray-400">Recommendation</span>
				<span class="font-semibold text-blue-600 dark:text-blue-400">HOLD</span>
			</div>
		</div>
	</div>
</div>
