<script lang="ts">
    import Menu from "./components/Menu.svelte"

    const talonData = JSON.parse(
        document.getElementById("talon-data").textContent
    )

</script>

{#if talonData}
    <Menu {talonData} />
{/if}
