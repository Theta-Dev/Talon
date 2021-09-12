<script lang="ts">
    import {selectTextOnFocus} from "../util/inputDirectives"
    import ClickOutside from "svelte-click-outside"

    import Icon from "./Icon.svelte"

    enum MenuState {
        Hidden,
        Closed,
        Open,
    }

    function showMenu(): void {
        menuState = MenuState.Closed;
    }

    function openSearch(): void {
        menuState = MenuState.Open;
        searchInput.focus()
    }

    function closeMenu(buttonClick: boolean): void {
        // Mobile device: hide open menu on outside click
        if (isMobile()) {
            if (menuState > MenuState.Hidden) {
                menuState = MenuState.Hidden
            }
        }
        // Desktop device: only hide menu with close button
        else {
            if (menuState > MenuState.Closed || (buttonClick && menuState > MenuState.Hidden)) {
                menuState--
            }
        }
    }

    function isMobile(): boolean {
        return window.innerWidth < 768
    }

    // let menuState = isMobile() ? MenuState.Hidden : MenuState.Closed
    let menuState = MenuState.Open
    let menuIcon: string
    let searchInput: HTMLInputElement

    $: menuIcon = menuState === MenuState.Open ? 'arrowRight' : 'close'
</script>

<ClickOutside on:clickoutside={() => closeMenu(false)}>
    <div class="wrapper"
         class:show={menuState === MenuState.Open}
         class:hide={menuState === MenuState.Hidden}
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
                    <li class="button-close" on:click={() => closeMenu(true)}>
                        <Icon iconName={menuIcon} size="25"
                              marginRight="1em"/>
                        <span class="text">Close sidebar</span>
                    </li>
                </ul>
            </div>
        </nav>
    </div>
    <div class="fab" class:hide={menuState > MenuState.Hidden}
         on:click={showMenu}>
        <Icon iconName="menu" size="25"/>
    </div>
</ClickOutside>

<style lang="sass">
    $color-base: #282d34
    $color-text: #ededed
    $color-text-active: $color-text
    $color-primary: #7935df

    $mobile: 768px

    =placeholder
        ::-webkit-input-placeholder
            @content
        ::-moz-placeholder
            @content
            opacity: 1
        ::-ms-input-placeholder
            @content

    =hideScrollbar
        -ms-overflow-style: none
        scrollbar-width: none

        &::webkit-scrollbar
            display: none

    =mobile
        @media screen and (max-width: $mobile - 1px)
            @content

    =desktop
        @media screen and (min-width: $mobile)
            @content

    *
        margin: 0
        padding: 0
        box-sizing: border-box

        color: $color-text

    .wrapper
        position: fixed
        top: 0
        right: 0
        width: 56px
        height: 100%

        padding: 1.25rem 0.5rem 0

        font-family: sans-serif
        font-size: 1rem

        background: $color-base

        transition: width 0.5s, padding 0.5s

        &.show
            width: 250px

        &.hide
            width: 0
            padding: 0

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

            &.active, &:hover, &:focus-within
                background: $color-primary
                color: $color-text-active

            .icon
                margin-right: 1rem

            .text
                font-weight: bold
                white-space: nowrap

    .nav-inner
        overflow-x: hidden
        overflow-y: auto

        +hideScrollbar

    .search
        input
            border: none
            background: none
            width: 100%

            font-family: sans-serif
            font-weight: bold
            font-size: 1rem

            outline: none

        +placeholder

        &.active, &:hover, &:focus-within
            input
                color: $color-text-active

            +placeholder
                color: $color-text-active

    .fab
        position: fixed
        bottom: 25px
        right: 25px

        height: 56px
        width: 56px

        display: flex
        align-items: center
        justify-content: center

        cursor: pointer
        border-radius: 50%

        background: $color-primary
        opacity: 1

        transition: bottom 0.5s, opacity 0.5s

        &.hide
            bottom: -56px
            opacity: 0

    .button-close
        +mobile
            display: none

    .icon > img
        height: 25px
        width: 25px
</style>
