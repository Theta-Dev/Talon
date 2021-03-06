/*
This configuration was taken from stefanonepa's project template
Copyright (c) 2021 stefanonepa, MIT License
https://github.com/stefanonepa/svelte-component-ts

This is a modified version of rollup-plugin-css-only
Copyright (c) 2016 Thomas Ghysels
https://github.com/thgh/rollup-plugin-css-only
 */

import {createFilter} from "@rollup/pluginutils"

function css(options) {
	if (options === void 0) options = {}

	var filter = createFilter(
		["**/*.css", "**/*.svelte", "**/*.js", "**/*.ts"],
		options.exclude
	)
	var isCssOnly = createFilter(options.include || ["**/*.css"], options.exclude)
	var styles = {}
	var order = []
	var dest = options.output
	var changes = 0

	return {
		name: "css",
		transform: function transform(code, id) {
			if (!isCssOnly(id)) {
				return
			}

			// Track the order that each stylesheet is imported.
			if (!order.includes(id)) {
				order.push(id)
			}

			// Keep track of every stylesheet
			// Check if it changed since last render
			if (styles[id] !== code && (styles[id] || code)) {
				styles[id] = code
				changes++
			}
			return ""
		},
		generateBundle: function generateBundle(opts, bundle) {
			// Combine all stylesheets, respecting import order
			var css = ""
			for (var x = 0; x < order.length; x++) {
				var id = order[x]
				css += styles[id] || ""
			}

			// Emit styles through callback
			if (typeof options.output === "function") {
				options.output(css, styles, bundle)
			}
		},
	}
}

export default css
