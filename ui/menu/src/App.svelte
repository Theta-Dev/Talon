<script lang="ts">
    import Menu from "./components/Menu.svelte"

    const talonData = JSON.parse(
        document.getElementById("talon-data").textContent
    )

</script>

<div class="talon-sidebar-container">
    {#if talonData}
        <Menu {talonData} />
    {/if}
</div>
