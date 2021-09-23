<script lang="ts">
    import {selectTextOnFocus} from "../util/inputDirectives"

    import MenuItem from "./MenuItem.svelte"
    import Icon from "./Icon.svelte"
    import type {Focusable} from "../util/types"

    let inputElm: HTMLInputElement

    export let active: boolean = false
    export let text: string = ""

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
        pointer-events: auto

        background: none
        border: none
        outline: none
        text-align: right
        width: 10em

        display: none

        .active
            display: flex

        &:focus
            opacity: 1
            border-bottom: solid #fff 2px
</style>

<MenuItem {active} on:click>
    <input
        placeholder="Search..."
        bind:this={inputElm}
        bind:value={text}
        on:focusout
        on:keyup
        use:selectTextOnFocus />
    <Icon iconName="search" size="40" scale="0.6" />
</MenuItem>
