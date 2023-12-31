import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

import { EndUrl } from './url_config'
const frontEndUrl  = EndUrl().frontEndUrl
const frontEndPost  = EndUrl().frontpost

export default defineConfig({
  plugins: [
    vue(),
  ],

  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server:{
   // host: '192.168.0.103'
    host: frontEndUrl,
    port:frontEndPost
  }
})
