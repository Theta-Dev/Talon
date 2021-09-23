<script lang="ts">
    import {Modals, closeModal} from "svelte-modals"

    import Menu from "./components/Menu.svelte"
    import type {TalonData} from "./util/types"

    const talonData: TalonData = JSON.parse(
        /* global document */
        document.getElementById("talon-data").textContent
    ) as TalonData

</script>

<style lang="sass">
    talon-div
        position: fixed
        top: 0
        bottom: 0
        right: 0
        left: 0
        background: rgba(0, 0, 0, 0.6)
</style>

<talon-sidebar>
    {#if talonData}
        <Menu {talonData} />
    {/if}

    <Modals>
        <talon-div slot="backdrop" on:click={closeModal} />
    </Modals>
</talon-sidebar>
