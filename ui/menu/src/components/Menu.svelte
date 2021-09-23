<script lang="ts">
    import {openModal} from "svelte-modals"

    import Icon from "./Icon.svelte"
    import MenuItem from "./MenuItem.svelte"
    import MenuItemPage from "./MenuItemPage.svelte"
    import InfoModal from "./InfoModal.svelte"
    import FloatingButton from "./FloatingButton.svelte"

    import type {Focusable, TalonData, TalonPage} from "../util/types"
    import {TalonVisibility} from "../util/types"
    import PageIcon from "./PageIcon.svelte"
    import MenuItemInput from "./MenuItemInput.svelte"

    function showSidebar(): void {
        sidebarShown = true
    }

    function hideSidebar(): void {
        sidebarShown = false
    }

    function isMobile(): boolean {
        /* global window */
        return window.innerWidth < 768
    }

    function openSearch(): void {
        searchOpen = true
        searchInput.focus()
    }

    function closeSearch() {
        searchOpen = false
        searchInput.blur()

        if (displayedPages.length === 0) searchText = ""
    }

    function searchKeypress(e: KeyboardEvent) {
        switch (e.key) {
            case "Enter":
                if (!searchText) {
                    closeSearch()
                } else if (displayedPages.length) {
                    window.location.href =
                        talonData.root_path + displayedPages[0].path
                } else {
                    closeSearch()
                }
                break
            case "Escape":
                closeSearch()
                break
        }
    }

    function openInfo() {
        openModal(InfoModal, {
            data: talonData,
        })
    }

    export let talonData: TalonData

    let sidebarShown = !isMobile()
    let searchInput: Focusable
    let searchOpen = false
    let searchText = ""

    let currentPage: TalonPage
    $: currentPage = talonData.pages[talonData.current_page]

    let displayedPages: TalonPage[]
    $: displayedPages = Object.entries(talonData.pages)
        .filter(([id, page]) => {
            if (id === talonData.current_page) return false

            if (searchText) {
                return (
                    page.visibility !== TalonVisibility.HIDDEN &&
                    page.name.toLowerCase().includes(searchText.toLowerCase())
                )
            }
            return page.visibility === TalonVisibility.FEATURED
        })
        .map(([, page]) => page)

</script>

<style lang="sass">
    @use "../style/values"
    @use "../style/mixin"

    talon-nav
        position: fixed
        top: 0
        right: 0
        height: 100%

        padding: 1em 0.4em

        display: flex
        flex-direction: column
        justify-content: space-between
        overflow: hidden

        >talon-div
            flex: 2 1 auto
            overflow-x: hidden
            overflow-y: auto

            &:first-child, &:last-child
                flex: 0 0 auto

            +mixin.hideScrollbar
</style>

{#if sidebarShown}
    <talon-nav class:talon-hide={!sidebarShown}>
        <talon-div>
            <MenuItemInput
                active={searchOpen || Boolean(searchText).valueOf()}
                on:click={openSearch}
                on:focusout={closeSearch}
                on:keyup={searchKeypress}
                bind:input={searchInput}
                bind:text={searchText} />
        </talon-div>
        <talon-div>
            {#each displayedPages as page, i}
                <MenuItemPage
                    {page}
                    rootPath={talonData.root_path}
                    active={searchOpen && searchText && i === 0} />
            {/each}
        </talon-div>
        <talon-div>
            {#if currentPage.source}
                <MenuItem
                    text="View source"
                    link={currentPage.source.url}
                    newTab={true}
                    privacy={true}>
                    <Icon
                        iconName={currentPage.source.type}
                        size={40}
                        scale={0.6} />
                </MenuItem>
            {/if}
            <MenuItem text="Info" on:click={openInfo}>
                <PageIcon page={currentPage} />
            </MenuItem>
            <MenuItem text="Hide sidebar" on:click={hideSidebar}>
                <Icon iconName="arrowRight" size={40} scale={0.6} />
            </MenuItem>
        </talon-div>
    </talon-nav>
{/if}

<FloatingButton hide={sidebarShown} on:click={showSidebar} />
