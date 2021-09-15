<script lang="ts">
    export let imageSrc: string
    export let size: number = 32
    export let scale: number = 1
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
    {/if}
</span>


<style lang="sass">
    img
        object-fit: cover
</style>
