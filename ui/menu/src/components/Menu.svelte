<script lang="ts">
    import {selectTextOnFocus} from "../util/inputDirectives"
    import ClickOutside from "svelte-click-outside"

    import Icon from "./Icon.svelte"

    enum MenuState {
        Hidden,
        Closed,
        Open,
    }

    function openSearch(): void {
        menuState = MenuState.Open;
        searchInput.focus()
    }

    function closeMenu(): void {
        if (menuState > MenuState.Closed) {
            menuState--
        }
    }

    function closeHideMenu(): void {
        if (menuState > MenuState.Hidden) {
            menuState--
        }
    }

    let menuState = MenuState.Closed
    let menuIcon: string
    let searchInput: HTMLInputElement

    $: menuIcon = menuState === MenuState.Open ? 'arrowRight' : 'close'
</script>

<ClickOutside on:clickoutside={closeMenu}>
    {#if menuState > MenuState.Hidden}
        <div class="wrapper"
             class:show={menuState === MenuState.Open}
        >
            <nav>
                <div class="nav-inner">
                    <ul>
                        <li class="search" on:click={openSearch}>
                            <Icon iconName="search" size="25"
                                  marginRight="1em"/>
                            <input
                                bind:this={searchInput}
                                use:selectTextOnFocus
                                placeholder="Search"
                            />
                        </li>
                    </ul>
                </div>

                <div class="nav-inner">
                    <ul>
                        <li>
                            <Icon iconName="github" size="25"
                                  marginRight="1em"/>
                            <span class="text">View source</span>
                        </li>
                        <li>
                            <Icon iconName="history" size="25"
                                  marginRight="1em"/>
                            <span class="text">Version history</span>
                        </li>
                        <li>
                            <Icon iconName="info" size="25"
                                  marginRight="1em"/>
                            <span class="text">Site info</span>
                        </li>
                        <li on:click={closeHideMenu}>
                            <Icon iconName={menuIcon} size="25"
                                  marginRight="1em"/>
                            <span class="text">Close sidebar</span>
                        </li>
                    </ul>
                </div>
            </nav>
        </div>
    {/if}
</ClickOutside>

<style lang="sass">
    @mixin placeholder
        ::-webkit-input-placeholder
            @content
        ::-moz-placeholder
            @content
            opacity: 1
        ::-ms-input-placeholder
            @content

    @mixin hideScrollbar
        -ms-overflow-style: none
        scrollbar-width: none

        &::webkit-scrollbar
            display: none

    $color-base: #1f242b
    $color-text: #ededed
    $color-text-active: $color-text
    $color-primary: #7935df

    *
        margin: 0
        padding: 0
        box-sizing: border-box

    .wrapper
        position: fixed
        top: 0
        right: 0
        width: 56px
        height: 100vh

        padding: 1.25rem 0.5rem 0

        font-family: sans-serif
        font-size: 1rem

        background: $color-base
        color: $color-text

        transition: width 0.5s

        &.show
            width: 250px

    nav
        height: 100%
        display: flex
        flex-direction: column
        justify-content: space-between
        overflow: hidden

        li, .search
            display: flex
            align-items: center

            cursor: pointer

            padding: 8px
            margin-bottom: 1rem
            border-radius: 0.5rem

            text-decoration: none

            color: $color-text

            &.active, &:hover, &:focus-within
                background: $color-primary
                color: $color-text-active

            .icon
                font-size: 20px
                margin-right: 1rem

            .text
                font-weight: bold
                white-space: nowrap

    .nav-inner
        overflow-x: hidden
        overflow-y: auto

        @include hideScrollbar

    .search
        input
            border: none
            background: none
            width: 100%

            font-family: sans-serif
            font-weight: bold
            font-size: 1rem

            color: $color-text

            outline: none

        @include placeholder
            color: $color-text

        &.active, &:hover, &:focus-within
            input
                color: $color-text-active

            @include placeholder
                color: $color-text-active
</style>
