import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import { federation } from "@module-federation/vite";

export default defineConfig({
  plugins: [
    vue(),
    federation({
      name: "advertising-widget-mf",
      filename: "remoteEntry.js",
      exposes: {
        "./AdvertisingWidget": "./src/AdvertisingWidget.ts",
        "./DiscountWidget": "./src/DiscountWidget.ts",
      },
      shared: ["vue"],
    }),
  ],
  server: {
    port: 5005,
    cors: true,
  },
  build: {
    target: "chrome89",
  },
});
