<script lang="ts">
	export let key = ""
	export let value = ""
	export let href: string = null

</script>

<style lang="sass">
	@use "../style/values"

	p
		margin: 0.3em 0
		display: flex

		a
			filter: none
			display: flex

		a:hover
			text-decoration: none
			filter: brightness(130%)

		span
			padding: 0.4em
			background-color: values.$color-base-1

		span:first-child
			font-weight: bold
			background-color: var(--talon-color)
</style>

<p>
	{#if href}
		<a {href}> <span>{key}</span> <span>{value}</span> </a>
	{:else}<span>{key}</span> <span>{value}</span>{/if}
</p>
