<script lang="ts">
    import icons from "../assets/icons"

    export let iconName: string
    export let color: string = "#fff"
    export let size: number = 32
    export let scale: number = 1
    export let dot: boolean = false

    let icon: [number, number, string]
    $: icon = icons[iconName] ?? [0, 0, ""]

</script>

<style lang="sass">
    @use "../style/values"

    .icon
        position: relative

    .dot
        width: 10px
        height: 10px
        border-radius: 50%
        position: absolute
        background: values.$color-primary-light
        bottom: 5%
        right: 5%
</style>

<span class="icon" style="width: {size}px; height: {size}px;">
    {#if icon}
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
    {/if}
    {#if dot}<span class="dot" />{/if}
</span>
