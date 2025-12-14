<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import * as d3 from 'd3';
	import type { StockData, ChartDimensions } from '$lib/types/stock';
	import { theme } from '$lib/stores/theme';

	interface Props {
		data: StockData[];
		symbol?: string;
		height?: number;
		width?: number;
	}

	let { data, symbol = 'STOCK', height = 400, width }: Props = $props();

	let chartContainer: HTMLDivElement;
	let svg: d3.Selection<SVGSVGElement, unknown, null, undefined>;
	let tooltip: d3.Selection<HTMLDivElement, unknown, null, undefined>;
	let crosshair: d3.Selection<SVGGElement, unknown, null, undefined>;
	let resizeObserver: ResizeObserver;
	let currentTheme: string = 'light';

	function getDimensions(): ChartDimensions {
		if (!chartContainer) {
			return {
				width: width || 800,
				height: height || 400,
				margin: { top: 20, right: 80, bottom: 40, left: 60 }
			};
		}

		const containerWidth = chartContainer.clientWidth;
		const containerHeight = chartContainer.clientHeight;
		return {
			width: containerWidth || width || 800,
			height: containerHeight || height || 400,
			margin: { top: 20, right: 80, bottom: 40, left: 60 }
		};
	}

	function getThemeColors() {
		return {
			green: currentTheme === 'dark' ? '#34d399' : '#10b981',
			red: currentTheme === 'dark' ? '#f87171' : '#ef4444',
			grid: currentTheme === 'dark' ? '#374151' : '#e5e7eb',
			text: currentTheme === 'dark' ? '#9ca3af' : '#6b7280',
			crosshair: currentTheme === 'dark' ? '#6b7280' : '#9ca3af',
			tooltipBg: currentTheme === 'dark' ? 'rgba(255, 255, 255, 0.9)' : 'rgba(0, 0, 0, 0.8)',
			tooltipText: currentTheme === 'dark' ? 'black' : 'white'
		};
	}

	onMount(() => {
		if (!chartContainer || !data || data.length === 0) return;

		// Initialize theme
		theme.subscribe((t) => {
			currentTheme = t;
			if (tooltip) {
				const colors = getThemeColors();
				tooltip.style('background', colors.tooltipBg).style('color', colors.tooltipText);
			}
			if (data && data.length > 0) {
				createChart();
			}
		});

		// Create tooltip
		const colors = getThemeColors();
		tooltip = d3
			.select(document.body)
			.append('div')
			.attr('class', 'chart-tooltip')
			.style('position', 'absolute')
			.style('visibility', 'hidden')
			.style('background', colors.tooltipBg)
			.style('color', colors.tooltipText)
			.style('padding', '8px')
			.style('border-radius', '4px')
			.style('font-size', '12px')
			.style('pointer-events', 'none')
			.style('z-index', '1000');

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
		if (tooltip) {
			tooltip.remove();
		}
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
			.attr('id', 'chart-clip')
			.append('rect')
			.attr('width', innerWidth)
			.attr('height', innerHeight);

		// Create scales
		const xScale = d3
			.scaleTime()
			.domain(d3.extent(data, (d) => d.date) as [Date, Date])
			.range([0, innerWidth]);

		const minLow = d3.min(data, (d: StockData) => d.low);
		const maxHigh = d3.max(data, (d: StockData) => d.high);

		const yScale = d3
			.scaleLinear()
			.domain([minLow! * 0.95, maxHigh! * 1.05] as [number, number])
			.range([innerHeight, 0]);

		// Create candlesticks
		const candleWidth = (innerWidth / data.length) * 0.6;
		const candleSpacing = innerWidth / data.length;

		// Draw wicks (high-low lines)
		g.selectAll('.wick')
			.data(data)
			.enter()
			.append('line')
			.attr('class', 'wick')
			.attr('clip-path', 'url(#chart-clip)')
			.attr('x1', (d: any) => xScale(d.date) + candleSpacing / 2)
			.attr('x2', (d: any) => xScale(d.date) + candleSpacing / 2)
			.attr('y1', (d: any) => yScale(d.high))
			.attr('y2', (d: any) => yScale(d.low))
			.attr('stroke', (d: any) => (d.close >= d.open ? colors.green : colors.red))
			.attr('stroke-width', 1);

		// Draw candle bodies
		g.selectAll('.candle')
			.data(data)
			.enter()
			.append('rect')
			.attr('class', 'candle')
			.attr('clip-path', 'url(#chart-clip)')
			.attr('x', (d: any) => xScale(d.date) + (candleSpacing - candleWidth) / 2)
			.attr('y', (d: any) => yScale(Math.max(d.open, d.close)))
			.attr('width', candleWidth)
			.attr('height', (d: any) => Math.abs(yScale(d.open) - yScale(d.close)))
			.attr('fill', (d: any) => (d.close >= d.open ? colors.green : colors.red))
			.attr('stroke', (d: any) => (d.close >= d.open ? colors.green : colors.red))
			.attr('stroke-width', 1)
			.on('mouseover', function (event: any, d: any) {
				showTooltip(event, d);
			})
			.on('mousemove', function (event: any) {
				moveTooltip(event);
			})
			.on('mouseout', function () {
				hideTooltip();
			});

		// Create axes
		const xAxis = d3
			.axisBottom(xScale)
			.ticks(Math.min(10, data.length))
			.tickFormat((d: any) => d3.timeFormat('%m/%d')(d));

		const yAxis = d3
			.axisRight(yScale)
			.ticks(8)
			.tickFormat((d: any) => d3.format('.2f')(d));

		g.append('g')
			.attr('class', 'x-axis')
			.attr('transform', `translate(0,${innerHeight})`)
			.call(xAxis);

		g.append('g')
			.attr('class', 'y-axis')
			.attr('transform', `translate(${innerWidth},0)`)
			.call(yAxis);

		// Style axes
		g.selectAll('.x-axis text, .y-axis text').style('font-size', '11px').style('fill', colors.text);

		g.selectAll('.x-axis line, .y-axis line, .x-axis path, .y-axis path')
			.style('stroke', colors.grid)
			.style('stroke-width', 1);

		// Create crosshair
		crosshair = g.append('g').attr('class', 'crosshair').style('display', 'none');

		crosshair
			.append('line')
			.attr('class', 'crosshair-x')
			.attr('x1', 0)
			.attr('x2', innerWidth)
			.attr('y1', 0)
			.attr('y2', 0)
			.attr('stroke', colors.crosshair)
			.attr('stroke-width', 1)
			.attr('stroke-dasharray', '2,2');

		crosshair
			.append('line')
			.attr('class', 'crosshair-y')
			.attr('x1', 0)
			.attr('x2', 0)
			.attr('y1', 0)
			.attr('y2', innerHeight)
			.attr('stroke', colors.crosshair)
			.attr('stroke-width', 1)
			.attr('stroke-dasharray', '2,2');

		// Add crosshair interaction
		svg
			.on('mousemove', function (event: any) {
				const [mouseX, mouseY] = d3.pointer(event);
				// Constrain crosshair to chart boundaries
				const constrainedX = Math.max(0, Math.min(mouseX - dimensions.margin.left, innerWidth));
				const constrainedY = Math.max(0, Math.min(mouseY - dimensions.margin.top, innerHeight));

				crosshair.style('display', null);
				crosshair.select('.crosshair-x').attr('y1', constrainedY).attr('y2', constrainedY);
				crosshair.select('.crosshair-y').attr('x1', constrainedX).attr('x2', constrainedX);
			})
			.on('mouseleave', function () {
				crosshair.style('display', 'none');
			});
	}

	function showTooltip(_event: MouseEvent, d: StockData) {
		const formatDate = d3.timeFormat('%Y-%m-%d');
		const formatPrice = d3.format('.2f');

		tooltip
			.html(
				`
			<div><strong>${symbol}</strong></div>
			<div>${formatDate(d.date)}</div>
			<div>O: ${formatPrice(d.open)}</div>
			<div>H: ${formatPrice(d.high)}</div>
			<div>L: ${formatPrice(d.low)}</div>
			<div>C: ${formatPrice(d.close)}</div>
			${d.volume ? `<div>Vol: ${d.volume.toLocaleString()}</div>` : ''}
		`
			)
			.style('visibility', 'visible');
	}

	function moveTooltip(event: MouseEvent) {
		tooltip.style('top', event.pageY - 10 + 'px').style('left', event.pageX + 10 + 'px');
	}

	function hideTooltip() {
		tooltip.style('visibility', 'hidden');
	}
</script>

<div
	class="w-full bg-white dark:bg-gray-800 rounded-lg shadow-lg p-4 transition-colors duration-200"
>
	<div class="mb-4">
		<h2 class="text-lg font-semibold text-gray-900 dark:text-gray-100">{symbol} Chart</h2>
	</div>
	<div bind:this={chartContainer} class="w-full" style="height: {height}px;"></div>
</div>
