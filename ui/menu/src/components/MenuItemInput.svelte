<script lang="ts">
    import {selectTextOnFocus} from "../util/inputDirectives"

    import MenuItem from "./MenuItem.svelte"
    import Icon from "./Icon.svelte"
    import type {Focusable} from "../util/types"

    let inputElm: HTMLInputElement

    export let active = false
    export let text = ""

    export const input: Focusable = {
        focus() {
            inputElm.focus()
        },
        blur() {
            inputElm.blur()
        },
    }

</script>

<style lang="sass">
    input
        background: none !important
        border: none !important
        outline: none !important
        text-align: right !important
        width: 10em !important

        display: none !important

        &:focus
            border-bottom: solid #fff 2px !important
</style>

<MenuItem {active} on:click>
    <input
        placeholder="Search..."
        bind:this={inputElm}
        bind:value={text}
        on:focusout
        on:keyup
        use:selectTextOnFocus />
    <Icon iconName="search" size={40} scale={0.6} />
</MenuItem>
