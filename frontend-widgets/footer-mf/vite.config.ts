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
      name: 'footer-mf',
      filename: 'remoteEntry.js',
      exposes: {
        './Footer': './src/footer/Footer.vue'
      },
      shared: {
        vue: { singleton: true },    
      },
    }),
  ],
  server: { port: 5002 }, 
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
})
