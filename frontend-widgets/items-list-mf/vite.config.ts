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
      name: 'items-list-mf',
      filename: 'remoteEntry.js',
      exposes: {
        './ItemsList': './src/items-list/ItemsList.vue'
      },
      shared: {
        vue: { singleton: true },    
      },
    }),
  ],
  server: { port: 5003 }, 
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
})
