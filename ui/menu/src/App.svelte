<script lang="ts">
    import {Modals, closeModal} from "svelte-modals"

    import Menu from "./components/Menu.svelte"
    import type {TalonData, TalonPage} from "./util/types"

    const talonData: TalonData = JSON.parse(
        /* global document */
        document.getElementById("talon-data").textContent
    ) as TalonData

    let currentPage: TalonPage
    $: currentPage = talonData.pages[talonData.current_page]

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

<talon-sidebar style="--talon-color: {currentPage.color}">
    {#if talonData}
        <Menu {talonData} />
    {/if}

    <Modals>
        <talon-div slot="backdrop" on:click={closeModal} />
    </Modals>
</talon-sidebar>
