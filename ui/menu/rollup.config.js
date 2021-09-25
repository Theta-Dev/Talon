/*
This configuration was taken from stefanonepa's project template
Copyright (c) 2021 stefanonepa, MIT License
https://github.com/stefanonepa/svelte-component-ts
 */

import svelte from "rollup-plugin-svelte"
import commonjs from "@rollup/plugin-commonjs"
import resolve from "@rollup/plugin-node-resolve"
import livereload from "rollup-plugin-livereload"
import {terser} from "rollup-plugin-terser"
import sveltePreprocess from "svelte-preprocess"
import typescript from "@rollup/plugin-typescript"
import css from ".rollup/css-only"
import {serve} from ".rollup/serve"
import replace from "@rollup/plugin-replace"
import babel from "@rollup/plugin-babel"
import includeSass from ".rollup/includeSass"

const production = !process.env.ROLLUP_WATCH
const extensions = [".svelte", ".ts", ".js", ".mjs"]

const bundleName = "bundle"
const bundleFile = `${bundleName}.js`
const bundleDir = production ? "dist" : "public"

const appFile = "src/App.svelte"
const styleFile = "src/style/main.sass"

function includeCss(styles, bundle) {
    const match = production
        ? `.shadowRoot.innerHTML="`
        : `.shadowRoot.innerHTML = "`

    const currentBundle = bundle[bundleFile]
    currentBundle.code = currentBundle.code.replace(
        match,
        `${match}<style>${styles}</style>`
    )
}

export default {
    input: "src/index.ts",
    output: [
        {
            sourcemap: !production,
            format: "iife",
            name: bundleName,
            file: `${bundleDir}/${bundleFile}`,
            plugins: [production && terser()],
        },
    ],
    plugins: [
        svelte({
            preprocess: sveltePreprocess({sourceMap: !production}),
            compilerOptions: {
                dev: !production,
                customElement: true,
            },
            emitCss: false,
            include: appFile,
        }),

        svelte({
            preprocess: sveltePreprocess({sourceMap: !production}),
            compilerOptions: {
                dev: !production,
            },
            emitCss: true,
            exclude: appFile,
        }),

        css({
            output(styles, styleNodes, bundle) {
                includeCss(styles, bundle)
            },
        }),

        includeSass({
            file: styleFile,
            outputStyle: "compressed",
            output: includeCss,
        }),

        resolve({
            browser: true,
            dedupe: ["svelte"],
            extensions,
        }),
        commonjs(),
        typescript({
            sourceMap: !production,
            inlineSources: !production,
        }),

        !production && serve(),

        !production && livereload(bundleDir),

        // add transition into shadow dom
        replace({
            ".ownerDocument": ".getRootNode()",
            delimiters: ["", ""],
            preventAssignment: true,
        }),
        replace({
            ".head.appendChild": ".appendChild",
            delimiters: ["", ""],
            preventAssignment: true,
        }),
        babel({
            extensions,
            exclude: "node_modules/**",
            plugins: ["@babel/plugin-proposal-class-properties"],
            presets: [
                [
                    "@babel/preset-env",
                    {
                        modules: false,
                        targets: {
                            esmodules: true,
                        },
                    },
                ],
                "@babel/preset-typescript",
            ],
            babelHelpers: "bundled",
        }),
    ],
    watch: {
        chokidar: true,
        clearScreen: false,
    },
    external: ["./src/style/test.css"],
}
