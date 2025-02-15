import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    host: "0.0.0.0",
    port: 3000,
    proxy: {
      "/api": {
        target: "https://hackathon-8.orb.ru/gis/",
        rewrite: (path) => path.replace(/^\/api/, ""),
        secure: false,
      },
    },
  },
  resolve: {
    alias: {
      app: "/src/app/",
      features: "/src/features/",
      pages: "/src/pages/",
      shared: "/src/shared/",
      "~@diplodoc": "/node_modules/@diplodoc/",
      url: "url",
    },
  },
});
