/** @type {import("snowpack").SnowpackUserConfig } */
export default {
    mount: {
        public: "/",
        src: "/",
    },
    plugins: ["@snowpack/plugin-svelte", "@snowpack/plugin-sass"],
    routes: [{match: "routes", src: ".*", dest: "/index.html"}],
    optimize: {
        bundle: true,
        minify: true,
    },
    packageOptions: {
        knownEntrypoints: ["svelte", "svelte/store"],
    },
    devOptions: {},
    buildOptions: {},
}
