import {defineConfig} from "vite";
import react from "@vitejs/plugin-react-swc";
import tailwindcss from "@tailwindcss/vite";
import path from "path";

export default defineConfig(() => ({
  plugins: [react(), tailwindcss()],
  build: {
    target: "baseline-widely-available",
    modulePreload: {
      polyfill: true,
    },
    outDir: "dist",
    minify: "esbuild",
    cssMinify: "esbuild",
    sourcemap: false,
    ssr: false,
    reportCompressedSize: true,
    rollupOptions: {
      output: {
        manualChunks(id: string) {
          if (id.includes("node_modules")) {
            return "vendor";
          }
          if (id.includes("shared/fe")) {
            return "shared-fe";
          }
        },
      },
    },
  },
  resolve: {
    alias: {
      "@shared/api-client": path.resolve(__dirname, "../shared/fe/api-client/dist/index"),
    },
  },
  server: {
    port: 35230,
  },
}));
