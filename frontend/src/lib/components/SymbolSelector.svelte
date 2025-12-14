<script lang="ts">
	interface Props {
		selectedSymbol: string;
		onSymbolChange: (symbol: string) => void;
	}

	let { selectedSymbol, onSymbolChange }: Props = $props();

	const popularStocks = [
		{ symbol: 'AAPL', name: 'Apple Inc.' },
		{ symbol: 'GOOGL', name: 'Alphabet Inc.' },
		{ symbol: 'MSFT', name: 'Microsoft Corporation' },
		{ symbol: 'TSLA', name: 'Tesla, Inc.' },
		{ symbol: 'AMZN', name: 'Amazon.com, Inc.' },
		{ symbol: 'META', name: 'Meta Platforms, Inc.' },
		{ symbol: 'NVDA', name: 'NVIDIA Corporation' },
		{ symbol: 'JPM', name: 'JPMorgan Chase & Co.' },
		{ symbol: 'V', name: 'Visa Inc.' },
		{ symbol: 'WMT', name: 'Walmart Inc.' }
	];

	let inputSymbol = $state('');
	let showSuggestions = $state(false);
	let filteredStocks = $derived(() => {
		if (inputSymbol) {
			return popularStocks.filter(
				(stock) =>
					stock.symbol.toLowerCase().includes(inputSymbol.toLowerCase()) ||
					stock.name.toLowerCase().includes(inputSymbol.toLowerCase())
			);
		} else {
			return popularStocks;
		}
	});

	// Initialize inputSymbol when selectedSymbol changes
	$effect(() => {
		inputSymbol = selectedSymbol;
	});

	function handleInput(event: Event) {
		const target = event.target as HTMLInputElement;
		inputSymbol = target.value.toUpperCase();
		showSuggestions = true;
	}

	function handleSymbolSelect(symbol: string) {
		inputSymbol = symbol;
		selectedSymbol = symbol;
		showSuggestions = false;
		onSymbolChange(symbol);
	}

	function handleSubmit(event: Event) {
		event.preventDefault();
		if (inputSymbol.trim()) {
			handleSymbolSelect(inputSymbol.trim());
		}
	}

	function handleBlur() {
		// Delay hiding suggestions to allow click events
		setTimeout(() => {
			showSuggestions = false;
		}, 200);
	}

	function handleFocus() {
		showSuggestions = true;
	}

	// Close suggestions when clicking outside
	function handleClickOutside(event: MouseEvent) {
		const target = event.target as HTMLElement;
		if (!target.closest('.symbol-selector')) {
			showSuggestions = false;
		}
	}
</script>

<svelte:window onclick={handleClickOutside} />

<div class="symbol-selector relative">
	<label for="symbol-input" class="block text-sm font-medium text-gray-700 mb-2">
		Stock Symbol
	</label>
	<form onsubmit={handleSubmit}>
		<div class="relative">
			<input
				id="symbol-input"
				type="text"
				bind:value={inputSymbol}
				oninput={handleInput}
				onfocus={handleFocus}
				onblur={handleBlur}
				placeholder="Enter stock symbol (e.g., AAPL)"
				class="w-full px-4 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500 pr-10"
			/>
			<div class="absolute inset-y-0 right-0 flex items-center pr-3 pointer-events-none">
				<svg class="w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"
					></path>
				</svg>
			</div>
		</div>
	</form>

	{#if showSuggestions && filteredStocks().length > 0}
		<div
			class="absolute z-10 w-full mt-1 bg-white border border-gray-300 rounded-md shadow-lg max-h-60 overflow-auto"
		>
			{#each filteredStocks() as stock}
				<button
					type="button"
					onclick={() => handleSymbolSelect(stock.symbol)}
					class="w-full px-4 py-2 text-left hover:bg-gray-100 focus:bg-gray-100 focus:outline-none border-b border-gray-100 last:border-b-0"
				>
					<div class="flex justify-between items-center">
						<span class="font-medium text-gray-900">{stock.symbol}</span>
						<span class="text-sm text-gray-500 truncate ml-2">{stock.name}</span>
					</div>
				</button>
			{/each}
		</div>
	{/if}
</div>
