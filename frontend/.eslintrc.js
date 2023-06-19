module.exports = {
  root: true,
  env: {
    browser: true,
    "vue/setup-compiler-macros": true,
  },
  extends: [
    "airbnb-base",
    "plugin:vue/vue3-recommended",
    "prettier",
    "plugin:prettier/recommended",
  ],
  parserOptions: {
    sourceType: "module",
  },
  rules: {
    "no-console": process.env.NODE_ENV === "production" ? "warn" : "off",
    "no-debugger": process.env.NODE_ENV === "production" ? "warn" : "off",
    "max-len": "off",
    "import/prefer-default-export": "off",
    "no-restricted-syntax": "off",
    "no-param-reassign": ["error", { props: false }],
    "vue/require-v-for-key": "off",
    "vue/valid-v-for": "off",
    "no-labels": ["error", { allowLoop: true }],
    "import/extensions": "off",
    "import/no-unresolved": "off",
  },
};
