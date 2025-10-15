import js from "@eslint/js";
import globals from "globals";
import reactHooks from "eslint-plugin-react-hooks";
import reactRefresh from "eslint-plugin-react-refresh";
import tsParser from "@typescript-eslint/parser";
import {defineConfig, globalIgnores} from "eslint/config";
import eslintConfigPrettier from "eslint-config-prettier/flat";
import tseslint from "typescript-eslint";

export default defineConfig([
  eslintConfigPrettier,
  tseslint.configs.recommended,
  reactHooks.configs.flat["recommended-latest"],
  reactRefresh.configs.recommended,
  reactRefresh.configs.vite,
  {
    files: ["**/*.{ts,tsx}"],
    ignores: ["dist/*"],
    plugins: {
      js,
    },
    rules: {
      // Custom overrides
      "@typescript-eslint/no-explicit-any": "off",
      "@typescript-eslint/no-unused-vars": "warn",
      "@/prefer-const": "warn",
      "react-hooks/exhaustive-deps": "off",
    },
    languageOptions: {
      parser: tsParser,
      ecmaVersion: 2020,
      globals: {
        React: true,
        ...globals.browser,
        ...globals.node,
        ...globals.nodeBuiltin,
      },
    },
  },
  globalIgnores(["dist/*"]),
]);
