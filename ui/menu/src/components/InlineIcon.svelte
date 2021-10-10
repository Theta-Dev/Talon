<script lang="ts">
	import icons from "../util/icons"

	export let iconName: string
	export let color = "#fff"
	export let size = 24

	let icon: [number, number, string]
	$: icon = icons[iconName] ?? [0, 0, ""]

</script>

<style lang="sass">
	@use "../style/values"

	span
		display: inline-block
</style>

<span class="icon">
	<svg
		aria-hidden="true"
		focusable="false"
		role="img"
		xmlns="http://www.w3.org/2000/svg"
		width={size}
		height={size}
		viewBox="0 0 {icon[0]} {icon[1]}">
		<path d={icon[2]} fill={color} />
	</svg>
</span>
