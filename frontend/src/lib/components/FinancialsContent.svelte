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
	}

	let { companyData }: Props = $props();
</script>

<div class="bg-white dark:bg-gray-800 rounded-lg shadow-lg p-6">
	<div class="grid grid-cols-1 md:grid-cols-2 gap-6">
		<div class="space-y-3">
			<div
				class="flex justify-between items-center py-2 border-b border-gray-100 dark:border-gray-700"
			>
				<span class="text-sm font-medium text-gray-600 dark:text-gray-400">Revenue (TTM)</span>
				<span class="font-semibold text-gray-900 dark:text-gray-100"
					>{formatNumber(companyData.revenueTTM)}</span
				>
			</div>
			<div
				class="flex justify-between items-center py-2 border-b border-gray-100 dark:border-gray-700"
			>
				<span class="text-sm font-medium text-gray-600 dark:text-gray-400">Profit Margin</span>
				<span class="font-semibold text-gray-900 dark:text-gray-100"
					>{(companyData.profitMargin * 100).toFixed(1)}%</span
				>
			</div>
			<div
				class="flex justify-between items-center py-2 border-b border-gray-100 dark:border-gray-700"
			>
				<span class="text-sm font-medium text-gray-600 dark:text-gray-400">Return on Assets</span>
				<span class="font-semibold text-gray-900 dark:text-gray-100"
					>{(companyData.roa * 100).toFixed(1)}%</span
				>
			</div>
			<div class="flex justify-between items-center py-2">
				<span class="text-sm font-medium text-gray-600 dark:text-gray-400">Book Value</span>
				<span class="font-semibold text-gray-900 dark:text-gray-100">${companyData.bookValue}</span>
			</div>
		</div>
		<div class="space-y-3">
			<div
				class="flex justify-between items-center py-2 border-b border-gray-100 dark:border-gray-700"
			>
				<span class="text-sm font-medium text-gray-600 dark:text-gray-400">Forward P/E</span>
				<span class="font-semibold text-gray-900 dark:text-gray-100">{companyData.forwardPE}</span>
			</div>
			<div
				class="flex justify-between items-center py-2 border-b border-gray-100 dark:border-gray-700"
			>
				<span class="text-sm font-medium text-gray-600 dark:text-gray-400">Return on Equity</span>
				<span class="font-semibold text-gray-900 dark:text-gray-100"
					>{(companyData.roe * 100).toFixed(1)}%</span
				>
			</div>
			<div
				class="flex justify-between items-center py-2 border-b border-gray-100 dark:border-gray-700"
			>
				<span class="text-sm font-medium text-gray-600 dark:text-gray-400">Price to Book</span>
				<span class="font-semibold text-gray-900 dark:text-gray-100">{companyData.priceToBook}</span
				>
			</div>
			<div class="flex justify-between items-center py-2">
				<span class="text-sm font-medium text-gray-600 dark:text-gray-400">Dividend/Share</span>
				<span class="font-semibold text-gray-900 dark:text-gray-100"
					>${companyData.dividendPerShare}</span
				>
			</div>
		</div>
	</div>
</div>
