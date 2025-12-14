<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import * as d3 from 'd3';
	import type { StockData, ChartDimensions } from '$lib/types/stock';
	import { theme } from '$lib/stores/theme';

	interface Props {
		data: StockData[];
		height?: number;
		width?: number;
	}

	let { data, height = 120, width }: Props = $props();

	let chartContainer: HTMLDivElement;
	let svg: d3.Selection<SVGSVGElement, unknown, null, undefined>;
	let resizeObserver: ResizeObserver;
	let currentTheme: string = 'light';

	function getDimensions(): ChartDimensions {
		if (!chartContainer) {
			return {
				width: width || 800,
				height: height || 120,
				margin: { top: 10, right: 80, bottom: 20, left: 60 }
			};
		}

		const containerWidth = chartContainer.clientWidth;
		const containerHeight = chartContainer.clientHeight;
		return {
			width: containerWidth || width || 800,
			height: containerHeight || height || 120,
			margin: { top: 10, right: 80, bottom: 20, left: 60 }
		};
	}

	function getThemeColors() {
		return {
			green: currentTheme === 'dark' ? '#34d399' : '#10b981',
			red: currentTheme === 'dark' ? '#f87171' : '#ef4444',
			grid: currentTheme === 'dark' ? '#374151' : '#e5e7eb',
			text: currentTheme === 'dark' ? '#9ca3af' : '#6b7280',
			barOpacity: currentTheme === 'dark' ? 0.7 : 0.8
		};
	}

	function formatVolume(volume: number): string {
		if (volume >= 1e9) {
			return (volume / 1e9).toFixed(1) + 'B';
		} else if (volume >= 1e6) {
			return (volume / 1e6).toFixed(1) + 'M';
		} else if (volume >= 1e3) {
			return (volume / 1e3).toFixed(1) + 'K';
		} else {
			return volume.toString();
		}
	}

	onMount(() => {
		if (!chartContainer || !data || data.length === 0) return;

		// Initialize theme
		theme.subscribe((t) => {
			currentTheme = t;
			if (data && data.length > 0) {
				createChart();
			}
		});

		// Set up resize observer
		resizeObserver = new ResizeObserver(() => {
			if (data && data.length > 0) {
				createChart();
			}
		});
		resizeObserver.observe(chartContainer);

		createChart();
	});

	onDestroy(() => {
		if (resizeObserver) {
			resizeObserver.disconnect();
		}
	});

	// Reactive effect for data changes
	$effect(() => {
		if (data && data.length > 0 && chartContainer) {
			createChart();
		}
	});

	function createChart() {
		const dimensions = getDimensions();
		const colors = getThemeColors();

		// Clear existing chart
		d3.select(chartContainer).selectAll('*').remove();

		const innerWidth = dimensions.width - dimensions.margin.left - dimensions.margin.right;
		const innerHeight = dimensions.height - dimensions.margin.top - dimensions.margin.bottom;

		// Create SVG
		svg = d3
			.select(chartContainer)
			.append('svg')
			.attr('width', dimensions.width)
			.attr('height', dimensions.height);

		// Create main group
		const g = svg
			.append('g')
			.attr('transform', `translate(${dimensions.margin.left},${dimensions.margin.top})`);

		// Add clipping path to prevent overflow
		g.append('clipPath')
			.attr('id', 'volume-clip')
			.append('rect')
			.attr('width', innerWidth)
			.attr('height', innerHeight);

		// Create scales
		const xScale = d3
			.scaleTime()
			.domain(d3.extent(data, (d) => d.date) as [Date, Date])
			.range([0, innerWidth]);

		const maxVolume = d3.max(data, (d: StockData) => d.volume || 0);

		const yScale = d3
			.scaleLinear()
			.domain([0, maxVolume! * 1.1] as [number, number])
			.range([innerHeight, 0]);

		// Calculate bar width
		const barWidth = (innerWidth / data.length) * 0.8;
		const barSpacing = innerWidth / data.length;

		// Draw volume bars
		g.selectAll('.volume-bar')
			.data(data)
			.enter()
			.append('rect')
			.attr('class', 'volume-bar')
			.attr('clip-path', 'url(#volume-clip)')
			.attr('x', (d: any) => xScale(d.date) + (barSpacing - barWidth) / 2)
			.attr('y', (d: any) => yScale(d.volume || 0))
			.attr('width', barWidth)
			.attr('height', (d: any) => innerHeight - yScale(d.volume || 0))
			.attr('fill', (d: any) => (d.close >= d.open ? colors.green : colors.red))
			.attr('opacity', colors.barOpacity);

		// Create axes
		const xAxis = d3
			.axisBottom(xScale)
			.ticks(Math.min(8, data.length))
			.tickFormat((d: any) => d3.timeFormat('%m/%d')(d));

		const yAxis = d3
			.axisRight(yScale)
			.ticks(4)
			.tickFormat((d: any) => formatVolume(d));

		g.append('g')
			.attr('class', 'x-axis')
			.attr('transform', `translate(0,${innerHeight})`)
			.call(xAxis);

		g.append('g')
			.attr('class', 'y-axis')
			.attr('transform', `translate(${innerWidth},0)`)
			.call(yAxis);

		// Style axes
		g.selectAll('.x-axis text, .y-axis text').style('font-size', '10px').style('fill', colors.text);

		g.selectAll('.x-axis line, .y-axis line, .x-axis path, .y-axis path')
			.style('stroke', colors.grid)
			.style('stroke-width', 1);
	}
</script>

<div
	class="w-full bg-white dark:bg-gray-800 rounded-lg shadow-lg p-3 transition-colors duration-200"
>
	<div class="mb-2">
		<h3 class="text-sm font-semibold text-gray-900 dark:text-gray-100">Volume</h3>
	</div>
	<div bind:this={chartContainer} class="w-full" style="height: {height}px;"></div>
</div>
