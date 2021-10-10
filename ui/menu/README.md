# Talon Menu

This is the sidebar menu that Talon injects into any page it displays.
It was created using [Svelte](https://svelte.deb).

## Available Scripts

### npm run dev

Runs the app in the development mode.
Open http://localhost:5000 to view it in the browser.

### npm run pc

(Pre-commit): Lint and format.

### npm run build

Builds the application and outputs a `dist/talon.js` file,
ready to be deployed!

The menu can be injected by adding these tags to the bottom of any
html document:

```html
<script id="talon-data" type="application/json">
  {
    "root_path": "/",
    "current_page": "1",
    "current_version": "5",
    ...
  }
</script>
<script src="talon.js"></script>
```
