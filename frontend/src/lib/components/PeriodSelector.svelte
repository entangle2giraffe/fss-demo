<script lang="ts">
	interface Props {
		selectedPeriod: string;
		onPeriodChange: (period: string) => void;
	}

	let { selectedPeriod, onPeriodChange }: Props = $props();

	const periods = [
		{ value: 'daily', label: 'Daily' },
		{ value: 'weekly', label: 'Weekly' },
		{ value: 'monthly', label: 'Monthly' }
	];

	function handlePeriodChange(event: Event) {
		const target = event.target as HTMLSelectElement;
		const newPeriod = target.value;
		selectedPeriod = newPeriod;
		onPeriodChange(newPeriod);
	}
</script>

<div class="period-selector">
	<label for="period-select" class="block text-sm font-medium text-gray-700 mb-2">
		Time Period
	</label>
	<select
		id="period-select"
		bind:value={selectedPeriod}
		onchange={handlePeriodChange}
		class="w-full px-4 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500 bg-white"
	>
		{#each periods as period}
			<option value={period.value}>{period.label}</option>
		{/each}
	</select>
</div>