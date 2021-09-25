<script lang="ts">
    import {Modals, closeModal} from "svelte-modals"

    import Menu from "./components/Menu.svelte"
    import type {TalonData, TalonPage} from "./util/types"

    const talonData: TalonData = JSON.parse(
        document.getElementById("talon-data").textContent
    ) as TalonData

    let currentPage: TalonPage
    $: currentPage = talonData.pages[talonData.current_page]

</script>

<style lang="sass">
    // Default theme
    .wrapper
        --talon-color: #7935df
</style>

<svelte:options tag="talon-sidebar" />
<div class="wrapper" style="--talon-color: {currentPage.color}">
    {#if talonData}
        <Menu {talonData} />
    {/if}

    <Modals>
        <div class="backdrop" slot="backdrop" on:click={closeModal} />
    </Modals>
</div>
