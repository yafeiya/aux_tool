import { fileURLToPath, URL } from 'node:url'
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import electron from 'vite-plugin-electron'
import proxy from 'vite-plugin-proxy'
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    electron({
      entry: 'electron/main.js'
  },
  )
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },

    

  })
