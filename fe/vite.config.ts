import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";
import tailwindcss from "@tailwindcss/vite";
// import {visualizer} from "rollup-plugin-visualizer";

export default defineConfig(() => ({
  plugins: [
    react(),
    tailwindcss(),
    // svgr({
    //   svgrOptions: {
    //     plugins: ["@svgr/plugin-svgo", "@svgr/plugin-jsx"],
    //     svgoConfig: {
    //       multipass: true,
    //       floatPrecision: 2,
    //     },
    //   },
    // }),
    // visualizer({
    //   emitFile: true,
    //   filename: "stats.html",
    // }),
  ],
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
        },
      },
    },
  },
  server: {
    port: 35230,
  },
}));
