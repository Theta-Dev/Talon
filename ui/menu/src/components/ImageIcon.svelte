<script lang="ts">
    export let imageSrc: string
    export let size = 32
    export let scale = 1
    export let alt = "??"
    export let color = "#4b228a"

</script>

<style lang="sass">
    @use "../style/values"

    .icon > span
        background-position: center
        background-repeat: no-repeat
        background-size: cover
        text-transform: uppercase
        border-radius: 50%
</style>

<span class="icon" style="width: {size}px; height: {size}px;">
    {#if imageSrc}
        <span
            style="width: {size * scale}px; height: {size * scale}px; background-image: url('{imageSrc}')" />
    {:else}
        <span
            style="width: {size * scale}px; height: {size * scale}px; font-size: {size * scale * 0.55}px; background: {color}">
            {alt}
        </span>
    {/if}
</span>
