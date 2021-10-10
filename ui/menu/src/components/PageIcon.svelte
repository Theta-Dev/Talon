<script lang="ts">
	import ImageIcon from "./ImageIcon.svelte"
	import Icon from "./Icon.svelte"
	import type {TalonPage} from "../util/types"

	export let page: TalonPage
	export let size = 40
	export let scale = 0.8

</script>

{#if page}
	<ImageIcon
		imageSrc={page.image}
		color={page.color}
		alt={page.name.substr(0, 2)}
		{size}
		{scale} />
{:else}
	<Icon iconName="question" {size} scale={scale * 0.75} />
{/if}
