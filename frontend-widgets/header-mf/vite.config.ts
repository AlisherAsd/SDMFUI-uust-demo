import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'
import  { federation }  from '@module-federation/vite'

// https://vite.dev/config/
export default defineConfig({
  plugins: [  
    vue(),
    vueDevTools(),
    federation({
      name: 'header-mf',
      filename: 'remoteEntry.js',
      exposes: {
        './Header': './src/Header.vue'
      },
      shared: {
        vue: { singleton: true },    
      },
    }),
  ],
  server: { port: 5001 }, 
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
})
