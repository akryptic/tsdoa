import path from "node:path";

import { defineConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";

// https://vite.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      "@wails": path.resolve(__dirname, "wailsjs"),
    },
  },
  plugins: [svelte()],
});
