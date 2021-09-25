module.exports = {
    extends: ["eslint:recommended"],
    env: {browser: true, es6: true, node: true},
    parserOptions: {
        sourceType: "module",
    },
    overrides: [
        {
            files: ["*.ts", "*.svelte"],
            extends: [
                "eslint:recommended",
                "plugin:@typescript-eslint/eslint-recommended",
                "plugin:@typescript-eslint/recommended",
            ],
            globals: {
                Atomics: "readonly",
                SharedArrayBuffer: "readonly",
            },
            parser: "@typescript-eslint/parser",
            parserOptions: {
                project: "./tsconfig.json",
            },
            plugins: ["@typescript-eslint"],
        },
        {
            files: ["*.svelte"],
            processor: "svelte3/svelte3",
            parserOptions: {
                extraFileExtensions: [".svelte"],
            },
            plugins: ["svelte3", "@typescript-eslint"],
            settings: {
                "svelte3/typescript": true,
                "svelte3/ignore-styles": () => true,
            },
        },
    ],
    rules: {},
    ignorePatterns: [".rollup/**", "public/**", "dist/**"],
}
