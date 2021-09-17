<script lang="ts">
    export let imageSrc: string
    export let size: number = 32
    export let scale: number = 1
    export let alt: string = "--"
</script>

<span
    class="icon"
    style="width: {size}px; height: {size}px;"
>
    {#if imageSrc}
        <img
            src={imageSrc}
            style="width: {size * scale}px; height: {size * scale}px"
        >
    {:else}
        <span
            class="placeholder"
            style="width: {size * scale}px; height: {size * scale}px"
        >
            {alt}
        </span>
    {/if}
</span>


<style lang="sass">
    @use "../style/values"

    img, .placeholder
        object-fit: cover
        display: flex
        align-items: center
        justify-content: center

    .placeholder
        font-size: 1.2em !important
        text-transform: uppercase

        border-radius: 50%
        background-color: values.$color-primary-dark
</style>
