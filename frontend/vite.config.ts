import tailwindcss from '@tailwindcss/vite';
import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

export default defineConfig({
  plugins: [
    react(),
    tailwindcss(),
  ],
  server: {
    proxy: {
      '/shorten': 'http://localhost:8080',
      '/links': 'http://localhost:8080',
      '/:key': 'http://localhost:8080',
    },
  },
});
