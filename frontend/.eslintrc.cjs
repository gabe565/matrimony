/* eslint-env node */
require("@rushstack/eslint-patch/modern-module-resolution");

module.exports = {
  root: true,
  env: {
    browser: true,
  },
  extends: [
    "plugin:vue/vue3-recommended",
    "eslint:recommended",
    "@vue/eslint-config-prettier/skip-formatting",
    "prettier",
    "plugin:prettier/recommended",
  ],
  parserOptions: {
    sourceType: "module",
  },
  rules: {
    "vue/require-v-for-key": "warn",
    "vue/valid-v-for": "warn",
    "no-unused-vars": [
      "error",
      { varsIgnorePattern: "^_", argsIgnorePattern: "^_" },
    ],
    "import/extensions": "off",
  },
};
