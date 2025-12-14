<script lang="ts">
	interface Props {
		companyData: {
			symbol: string;
			name: string;
			exchange: string;
			sector: string;
		};
		currentPrice: number;
		priceChange: number;
		priceChangePercent: number;
	}

	let { companyData, currentPrice, priceChange, priceChangePercent }: Props = $props();
</script>

<div class="bg-white dark:bg-gray-800 rounded-lg shadow-lg p-6 mb-4">
	<div class="flex flex-col lg:flex-row lg:items-center lg:justify-between gap-4">
		<!-- Company Info -->
		<div class="company-info">
			<div class="flex items-center space-x-4 mb-2 lg:mb-0">
				<h1 class="text-3xl font-bold text-gray-900 dark:text-gray-100">
					{companyData.symbol}
				</h1>
				<div class="border-l border-gray-300 dark:border-gray-600 pl-4">
					<p class="text-lg font-medium text-gray-900 dark:text-gray-100">
						{companyData.name}
					</p>
					<p class="text-sm text-gray-600 dark:text-gray-400">
						{companyData.exchange} • {companyData.sector}
					</p>
				</div>
			</div>
		</div>

		<!-- Price Information -->
		<div class="price-info text-center lg:text-right">
			<div class="text-4xl font-bold text-gray-900 dark:text-gray-100 mb-1">
				${currentPrice.toFixed(2)}
			</div>
			<div
				class="text-lg font-medium {priceChange >= 0
					? 'text-green-600 dark:text-green-400'
					: 'text-red-600 dark:text-red-400'}"
			>
				{priceChange >= 0 ? '+' : ''}{priceChange.toFixed(2)} ({priceChangePercent >= 0
					? '+'
					: ''}{priceChangePercent.toFixed(2)}%)
			</div>
			<div class="text-sm text-gray-600 dark:text-gray-400 mt-1">
				Last updated: {new Date().toLocaleTimeString()}
			</div>
		</div>
	</div>
</div>
