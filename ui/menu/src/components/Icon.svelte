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

    .icon > span
        background: values.$var-primary
        border-radius: 50%

        &.transparent
            background: none
</style>

<span class="icon" style="width: {size}px; height: {size}px;">
    {#if icon}
        <span
            style="width: {(size * scale * 4) / 3}px; height: {(size * scale * 4) / 3}px"
            class:transparent>
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
        </span>
    {/if}
</span>
