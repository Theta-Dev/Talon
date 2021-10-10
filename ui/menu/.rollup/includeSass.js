/*
This is a modified version of rollup-plugin-scss
Copyright (c) 2016 Thomas Ghysels
https://github.com/thgh/rollup-plugin-scss
 */

function includeSass(options = {}) {
	let dest = typeof options.output === "string" ? options.output : null
	const prefix = options.prefix ? options.prefix + "\n" : ""
	let includePaths = options.includePaths || ["node_modules/"]
	includePaths.push(process.cwd())

	const compileToCSS = async function (file) {
		// Compile SASS to CSS
		if (file.length) {
			includePaths = includePaths.filter((v, i, a) => a.indexOf(v) === i)
			try {
				const sass = options.sass || loadSassLibrary()
				const render = sass.renderSync(
					Object.assign(
						{
							file,
							outFile: dest,
							includePaths,
							importer: (url, prev, done) => {
								/* If a path begins with `.`, then it's a local import and this
								 * importer cannot handle it. This check covers both `.` and
								 * `..`.
								 *
								 * Additionally, if an import path begins with `url` or `http`,
								 * then it's a remote import, this importer also cannot handle
								 * that. */
								if (
									url.startsWith(".") ||
									url.startsWith("url") ||
									url.startsWith("http")
								) {
									/* The importer returns `null` to defer processing the import
									 * back to the sass compiler. */
									return null
								}
								/* If the requested path begins with a `~`, we remove it. This
								 * character is used by webpack-contrib's sass-loader to
								 * indicate the import is from the node_modules folder. Since
								 * this is so standard in the JS world, the importer supports
								 * it, by removing it and ignoring it. */
								const cleanUrl = url.startsWith("~")
									? url.replace("~", "")
									: url
								/* Now, the importer uses `require.resolve()` to attempt
								 * to resolve the path to the requested file. In the case
								 * of a standard node_modules project, this will use Node's
								 * `require.resolve()`. In the case of a Plug 'n Play project,
								 * this will use the `require.resolve()` provided by the
								 * package manager.
								 *
								 * This statement is surrounded by a try/catch block because
								 * if Node or the package manager cannot resolve the requested
								 * file, they will throw an error, so the importer needs to
								 * defer to sass, by returning `null`.
								 *
								 * The paths property tells `require.resolve()` where to begin
								 * resolution (i.e. who is requesting the file). */
								try {
									const resolved = require.resolve(cleanUrl, {
										paths: [prefix + scss],
									})
									/* Since `require.resolve()` will throw an error if a file
									 * doesn't exist. It's safe to assume the file exists and
									 * pass it off to the sass compiler. */
									return {file: resolved}
								} catch (e) {
									/* Just because `require.resolve()` couldn't find the file
									 * doesn't mean it doesn't exist. It may still be a local
									 * import that just doesn't list a relative path, so defer
									 * processing back to sass by returning `null` */
									return null
								}
							},
						},
						options
					)
				)
				const css = render.css.toString()
				const map = render.map ? render.map.toString() : ""
				// Possibly process CSS (e.g. by PostCSS)
				if (typeof options.processor === "function") {
					const result = await options.processor(css, map, styles)
					// TODO: figure out how to check for
					// @ts-ignore
					const postcss = result
					// PostCSS support
					if (typeof postcss.process === "function") {
						return Promise.resolve(
							postcss.process(css, {
								from: undefined,
								to: dest,
								map: map ? {prev: map, inline: false} : null,
							})
						)
					}
					// @ts-ignore
					return stringToCSS(result)
				}
				return {css, map}
			} catch (e) {
				if (options.failOnError) {
					throw e
				}
				console.log()
				console.log(red("Error:\n\t" + e.message))
				if (e.message.includes("Invalid CSS")) {
					console.log(green("Solution:\n\t" + "fix your Sass code"))
					console.log("Line:   " + e.line)
					console.log("Column: " + e.column)
				}
				if (e.message.includes("sass") && e.message.includes("find module")) {
					console.log(green("Solution:\n\t" + "npm install --save-dev sass"))
				}
				if (e.message.includes("node-sass") && e.message.includes("bindings")) {
					console.log(
						green("Solution:\n\t" + "npm rebuild node-sass --force")
					)
				}
				console.log()
			}
		}
		return {css: "", map: ""}
	}

	return {
		name: "includeSass",

		buildStart() {
			this.addWatchFile(options.file)
		},

		async generateBundle(opts, bundle) {
			const compiled = await compileToCSS(options.file)
			if (typeof compiled !== "object" || typeof compiled.css !== "string") {
				return
			}

			// Emit styles through callback
			if (typeof options.output === "function") {
				options.output(compiled.css, bundle)
			}
		},
	}
}

function loadSassLibrary() {
	try {
		return require("sass")
	} catch (e) {
		return require("node-sass")
	}
}
function stringToCSS(input) {
	if (typeof input === "string") {
		return {css: input, map: ""}
	}
	return input
}
function red(text) {
	return "\x1b[1m\x1b[31m" + text + "\x1b[0m"
}
function green(text) {
	return "\x1b[1m\x1b[32m" + text + "\x1b[0m"
}

export default includeSass
