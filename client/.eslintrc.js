module.exports = {
  env: {
    browser: true,
    es2021: true,
    node: true,
    "react-native/react-native": true,
  },
  extends: [
    "eslint:recommended",
    "plugin:react/recommended",
    "plugin:@typescript-eslint/recommended",
    "plugin:react-native/all",
  ],
  parser: "@babel/eslint-parser",
  parserOptions: {
    requireConfigFile: false,
    ecmaFeatures: {
      jsx: true,
    },
    ecmaVersion: 12,
    sourceType: "module",
  },
  plugins: ["react", "@typescript-eslint", "react-native"],
  rules: {
    "no-console": "warn",
    "no-unused-vars": "warn",
    quotes: ["error", "single"],
    semi: ["error", "always"],
    indent: ["error", 2],
    "react-native/no-inline-styles": "warn",
    "react-native/no-raw-text": [
      "warn",
      {
        skip: ["Text"],
      },
    ],
  },
  settings: {
    react: {
      version: "detect",
    },
  },
};
