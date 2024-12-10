import globals from "globals";
import pluginJs from "@eslint/js";
import tseslint from "typescript-eslint";
import stylistic from "@stylistic/eslint-plugin";

/** @type {import('eslint').Linter.Config[]} */
export default [
  { files: ["**/*.{js,mjs,cjs,ts}"] },
  { files: ["**/*.js"], languageOptions: { sourceType: "script" } },
  { languageOptions: { globals: globals.browser } },
  { ignores: ["**/*.js"] },
  {
    plugins: {
      '@stylistic': stylistic
    }
  },
  pluginJs.configs.recommended,
  ...tseslint.configs.recommended,
  {
    rules: {
      "@stylistic/semi": ["error"],
      "@stylistic/semi-style": ["error"],
      "@stylistic/indent": ["warn", 2],
      "@stylistic/no-extra-semi": ["error"],
    }
  }
];