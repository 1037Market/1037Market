import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
  ],
  server: {
    host: '0.0.0.0',
    port: 14685
  },
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  compilerOptions: {
    isCustomElement: tagName => {
      return tagName === 'vue-advanced-chat' || tagName === 'emoji-picker'
    }
  },
  publicDir: '/1037Market/',
  publicPath: '/1037Market/',
  outputDir: '1037Market',
  assetsDir: 'assets',
  base: '/1037Market/',
  baseUrl: '/1037Market/'


})
