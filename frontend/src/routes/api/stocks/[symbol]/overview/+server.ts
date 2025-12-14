import { json, type RequestHandler } from '@sveltejs/kit';
import { env } from '$env/dynamic/private';

const GO_BASE_URL = env.GO_STOCKS_URL ?? 'http://localhost:8080';
const API_KEY = env.GO_API_KEY ?? '';

export const GET: RequestHandler = async ({ params }) => {
  const { symbol } = params;
  const targetUrl = `${GO_BASE_URL}/api/stocks/${symbol}/overview`;

  const headers: Record<string, string> = { 'Content-Type': 'application/json' };
  if (API_KEY) headers['X-API-Key'] = API_KEY;

  try {
    const resp = await fetch(targetUrl, { method: 'GET', headers });
    if (!resp.ok) {
      const text = await resp.text();
      console.error('Upstream overview error', { status: resp.status, body: text });
      return json({ error: text || resp.statusText }, { status: resp.status });
    }

    const overview = await resp.json();
    return json(overview);
  } catch (err: any) {
    console.error('Proxy error fetching overview:', err);
    return json({ error: 'Failed to fetch company overview' }, { status: 502 });
  }
};
