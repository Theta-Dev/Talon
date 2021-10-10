<script lang="ts">
	import Icon from "./Icon.svelte"

	export let hide = false

</script>

<style lang="sass">
	@use "../style/values"

	button
		position: fixed
		bottom: 25px
		right: 25px

		height: 56px
		width: 56px

		display: flex
		align-items: center
		justify-content: center

		border-radius: 50%
		background: values.$var-primary
		border: none

		opacity: 1
		transition: bottom 0.5s, opacity 0.5s

		&.hide
			bottom: -56px
			opacity: 0
</style>

<button class:hide on:click>
	<Icon iconName="menu" size={25} />
</button>
