<script lang="ts">
	interface Props {
		activeTab: string;
		onTabChange: (tab: string) => void;
	}

	let { activeTab, onTabChange }: Props = $props();

	function handleTabClick(tab: string) {
		onTabChange(tab);
	}
</script>

<div class="bg-white dark:bg-gray-800 rounded-lg shadow-lg p-4 mb-4">
	<nav class="flex flex-col sm:flex-row space-y-2 sm:space-y-0 sm:space-x-8">
		<button
			class="tab-button {activeTab === 'overview' ? 'active' : ''}"
			onclick={() => handleTabClick('overview')}
		>
			Overview
		</button>
		<button
			class="tab-button {activeTab === 'financials' ? 'active' : ''}"
			onclick={() => handleTabClick('financials')}
		>
			Financials
		</button>
		<button
			class="tab-button {activeTab === 'analysis' ? 'active' : ''}"
			onclick={() => handleTabClick('analysis')}
		>
			Analysis
		</button>
	</nav>
</div>

<style>
	.tab-button {
		padding-top: 0.75rem;
		padding-bottom: 0.75rem;
		padding-left: 1rem;
		padding-right: 1rem;
		border-bottom-width: 2px;
		font-size: 0.875rem;
		font-weight: 500;
		transition: all 0.2s;
		border-radius: 0.5rem 0.5rem 0 0;
	}

	.tab-button.active {
		border-color: rgb(59 130 246);
		color: rgb(37 99 235);
		background-color: rgb(239 246 255);
	}

	.tab-button:not(.active) {
		border-color: transparent;
		color: rgb(107 114 128);
	}

	.tab-button:hover:not(.active) {
		border-color: rgb(209 213 219);
		color: rgb(55 65 81);
		background-color: rgb(249 250 251);
	}

	@media (prefers-color-scheme: dark) {
		.tab-button.active {
			color: rgb(147 197 253);
			background-color: rgba(59 130 246 0.2);
		}

		.tab-button:not(.active) {
			color: rgb(156 163 175);
		}

		.tab-button:hover:not(.active) {
			border-color: rgb(75 85 99);
			color: rgb(209 213 219);
			background-color: rgba(55 65 81 0.5);
		}
	}
</style>
