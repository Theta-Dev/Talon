<script lang="ts">
    import {closeModal, Modals} from "svelte-modals"

    import Menu from "./components/Menu.svelte"
    import {currentPage, isPresent} from "./util/talonData"

</script>

<style lang="sass">
    // Default theme
    .wrapper
        --talon-color: #7935df

    .backdrop
        position: fixed
        top: 0
        bottom: 0
        right: 0
        left: 0
        background: rgba(0, 0, 0, 0.6)
</style>

<div class="wrapper" style={`--talon-color: ${currentPage.color}`}>
    {#if isPresent}
        <Menu />
    {/if}

    <Modals>
        <div class="backdrop" slot="backdrop" on:click={closeModal} />
    </Modals>
</div>
