import {createFilter} from "@rollup/pluginutils"
import {Minifier} from "minify-html-stream"
import {Readable} from "stream"

function htmlMinifier(options) {
    const filter = createFilter(["**/*.svelte"], options.exclude)

    return {
        name: "htmlMinifier",
        async transform(code, id) {
            if (!filter(id)) return code

            let stream = new Readable()
            stream.push(code)
            stream.push(null)
            stream = stream.pipe(new Minifier(options.options))

            const chunks = []

            for await (let chunk of stream) {
                chunks.push(chunk)
            }

            const buffer = Buffer.concat(chunks)

            return buffer.toString("utf-8")
        },
    }
}

export default htmlMinifier
