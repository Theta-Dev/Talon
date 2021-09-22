<script lang="ts">
    export let imageSrc: string
    export let size: number = 32
    export let scale: number = 1
    export let alt: string = "--"
    export let color: string = "#4b228a"

</script>

<style lang="sass">
    @use "../style/values"

    img, .talon-placeholder
        object-fit: cover
        display: flex
        align-items: center
        justify-content: center

    .talon-placeholder
        text-transform: uppercase

        border-radius: 50%
</style>

<span class="talon-icon" style="width: {size}px; height: {size}px;">
    {#if imageSrc}
        <img
            src={imageSrc}
            {alt}
            style="width: {size * scale}px; height: {size * scale}px" />
    {:else}
        <span
            class="talon-placeholder"
            style="width: {size * scale}px; height: {size * scale}px; font-size: {size * scale * 0.55}px; background: {color}">
            {alt}
        </span>
    {/if}
</span>
