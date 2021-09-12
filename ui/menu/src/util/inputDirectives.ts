/** Selects the text inside a text node when the node is focused */
export function selectTextOnFocus(node: HTMLInputElement) {
  const handleFocus = (event) => {
    node && typeof node.select === "function" && node.select();
  };

  node.addEventListener("focus", handleFocus);

  return {
    destroy() {
      node.removeEventListener("focus", handleFocus);
    },
  };
}
