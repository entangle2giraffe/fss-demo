export default function formatNumber(num: number): string {
	if (num >= 1e12) {
		return (num / 1e12).toFixed(1) + 'T'; // Trillion
	} else if (num >= 1e9) {
		return (num / 1e9).toFixed(1) + 'B'; // Billion
	} else if (num >= 1e6) {
		return (num / 1e6).toFixed(1) + 'M'; // Million
	} else {
		return num.toString(); // Return as is if less than a million
	}
}
