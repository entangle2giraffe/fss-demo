import { json, type RequestHandler } from '@sveltejs/kit';
import { env } from '$env/dynamic/private';

// Server-only env vars (dynamic/private avoids compile-time export issues)
const GO_BASE_URL = env.GO_STOCKS_URL ?? 'http://localhost:8080';
const API_KEY = env.GO_API_KEY ?? '';

function normalizeToArray(raw: any) {
  if (!raw) return [];
  if (Array.isArray(raw)) return raw;

  // Handle Go backend response format (entities.StockTimeSeries)
  if (raw.data && typeof raw.data === 'object' && !Array.isArray(raw.data)) {
    const arr = Object.entries(raw.data).map(([dateStr, stockData]: [string, any]) => ({
      date: new Date(dateStr),
      open: stockData.open ?? 0,
      high: stockData.high ?? 0,
      low: stockData.low ?? 0,
      close: stockData.close ?? 0,
      volume: stockData.volume ?? 0
    }));
    // Return newest-first to reduce client sorting work; loader will sort as needed
    return arr.sort((a: any, b: any) => b.date.getTime() - a.date.getTime());
  }

  // Handle direct Alpha Vantage response format (if passed through)
  if (raw['Time Series (Daily)'] || raw['Weekly Time Series'] || raw['Monthly Time Series']) {
    let timeSeriesKey = '';
    for (const key in raw) {
      if (key.includes('Time Series')) {
        timeSeriesKey = key;
        break;
      }
    }
    
    if (timeSeriesKey && raw[timeSeriesKey]) {
      const arr = Object.entries(raw[timeSeriesKey]).map(([dateStr, data]: [string, any]) => ({
        date: new Date(dateStr),
        open: parseFloat(data['1. open'] ?? '0'),
        high: parseFloat(data['2. high'] ?? '0'),
        low: parseFloat(data['3. low'] ?? '0'),
        close: parseFloat(data['4. close'] ?? '0'),
        volume: parseInt(data['5. volume'] ?? '0', 10)
      }));
      return arr.sort((a: any, b: any) => a.date.getTime() - b.date.getTime());
    }
  }

  // Fallback for other formats
  const values = Object.values(raw);
  if (values.length && typeof values[0] === 'object') return values;

  return [];
}

export const GET: RequestHandler = async ({ params, url }) => {
  const { symbol } = params;
  const targetUrl = `${GO_BASE_URL}/api/stocks/${symbol}${url.search}`;

  const headers: Record<string, string> = { 'Content-Type': 'application/json' };
  // Use the exact header name the Go middleware checks for
  if (API_KEY) headers['X-API-Key'] = API_KEY;

  try {
    const resp = await fetch(targetUrl, { method: 'GET', headers });
    if (!resp.ok) {
      const text = await resp.text();
      console.error('Upstream response error', { status: resp.status, body: text });
      // Return upstream body for easier debugging in the browser (server-only)
      return json({ error: text || resp.statusText }, { status: resp.status });
    }

    const raw = await resp.json();
    const arr = normalizeToArray(raw);
    return json({
      symbol: raw.symbol ?? symbol,
      data: arr
    });
  } catch (err: any) {
    console.error('Proxy error fetching stocks:', err);
    return json({ error: 'Failed to fetch stock data' }, { status: 502 });
  }
};