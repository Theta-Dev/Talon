<script lang="ts">
    import {Modals, closeModal} from "svelte-modals"

    import Menu from "./components/Menu.svelte"

    const talonData = JSON.parse(
        document.getElementById("talon-data").textContent
    )

</script>

<style lang="sass">
    div
        position: fixed
        top: 0
        bottom: 0
        right: 0
        left: 0
        background: rgba(0, 0, 0, 0.6)
</style>

<talon-sidebar id="talon-sidebar">
    {#if talonData}
        <Menu {talonData} />
    {/if}

    <Modals>
        <div slot="backdrop" on:click={closeModal} />
    </Modals>
</talon-sidebar>
