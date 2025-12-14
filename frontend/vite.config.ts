import tailwindcss from '@tailwindcss/vite';
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
    plugins: [
        tailwindcss(),
        sveltekit(),
    ],
    server: {
        host: '0.0.0.0',
        port: 5173,
        cors: true,
        
        // HTTPS HMR Configuration
        hmr: {
            // WebSocket configuration for HTTPS
            port: 5173,
            host: 'localhost',
            protocol: 'wss',
            
            // Client configuration for HTTPS
            clientPort: 5173,
            overlay: true
        },
        
        // Development proxy (fallback, Caddy handles primary routing)
        proxy: {
            '/api': {
                target: 'http://stock-ingestor:8080',
                changeOrigin: true,
                secure: false,
                rewrite: (path) => path.replace(/^\/api/, '')
            }
        }
    },
    
    preview: {
        host: '0.0.0.0',
        port: 4173,
        cors: true
    },
    
    // Build optimization
    build: {
        sourcemap: true,
        rollupOptions: {
            output: {
                manualChunks: {
                    vendor: ['d3'],
                    charts: ['@types/d3']
                }
            }
        }
    }
});
