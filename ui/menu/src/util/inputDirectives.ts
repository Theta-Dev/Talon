/** Selects the text inside a text node when the node is focused */
import type {SvelteActionRes} from "./types"

export function selectTextOnFocus(node: HTMLInputElement): SvelteActionRes {
	const handleFocus = () => {
		node && typeof node.select === "function" && node.select()
	}

	node.addEventListener("focus", handleFocus)

	return {
		destroy() {
			node.removeEventListener("focus", handleFocus)
		},
	}
}
