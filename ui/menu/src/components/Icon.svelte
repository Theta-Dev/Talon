<script lang="ts">
    import icons from "../assets/icons"

    export let iconName: string
    export let color = "#fff"
    export let size = 32
    export let scale = 1
    export let transparent = false

    let icon: [number, number, string]
    $: icon = icons[iconName] ?? [0, 0, ""]

</script>

<style lang="sass">
    @use "../style/values"

    talon-span
        display: flex
        align-items: center
        justify-content: center

        >talon-span
            background: values.$color-primary
            border-radius: 50%

            &.talon-transparent
                background: none
</style>

<talon-span style="width: {size}px; height: {size}px;">
    {#if icon}
        <talon-span
            style="width: {(size * scale * 4) / 3}px; height: {(size * scale * 4) / 3}px"
            class:talon-transparent={transparent}>
            <svg
                aria-hidden="true"
                focusable="false"
                role="img"
                xmlns="http://www.w3.org/2000/svg"
                width={size * scale}
                height={size * scale}
                viewBox="0 0 {icon[0]} {icon[1]}">
                <path d={icon[2]} fill={color} />
            </svg>
        </talon-span>
    {/if}
</talon-span>
