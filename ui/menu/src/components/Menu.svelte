<script lang="ts">
    import {closeModal, Modals, openModal} from "svelte-modals"
    import {selectTextOnFocus} from "../util/inputDirectives"

    import Icon from "./Icon.svelte"
    import MenuItem from "./MenuItem.svelte"
    import InfoModal from "./InfoModal.svelte"
    import FloatingButton from "./FloatingButton.svelte"

    import type {TalonData, TalonPage} from "../util/types"
    import {TalonVisibility} from "../util/types"
    import PageIcon from "./PageIcon.svelte"

    function showSidebar(): void {
        sidebarShown = true
    }

    function hideSidebar(): void {
        sidebarShown = false
    }

    function isMobile(): boolean {
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
        console.log(e.key)
        switch (e.key) {
            case "Enter":
                if (!searchText) {
                    closeSearch()
                } else if (displayedPages) {
                    window.location =
                        talonData.root_path + displayedPages[0].path
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

    let sidebarShown: boolean = !isMobile()
    let searchInput: HTMLInputElement
    let searchOpen: boolean = false
    let searchText: string = ""

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

    .backdrop
        position: fixed
        top: 0
        bottom: 0
        right: 0
        left: 0
        background: rgba(0, 0, 0, 0.6)
</style>

<div class="talon-wrapper" class:talon-hide={!sidebarShown}>
    <div class="talon-nav-inner" style="flex: 0 0 auto">
        <div
            class="talon-item"
            class:active={searchOpen || searchText}
            on:click={openSearch}>
            <span class="talon-text" />
            <input
                placeholder="Search..."
                bind:this={searchInput}
                bind:value={searchText}
                on:focusout={closeSearch}
                on:keyup={searchKeypress}
                use:selectTextOnFocus />
            <Icon iconName="search" size="40" scale="0.6" />
        </div>
    </div>
    <div class="talon-nav-inner" style="flex: 2 1 auto">
        {#each displayedPages as page, i}
            <MenuItem
                {page}
                rootPath={talonData.root_path}
                active={searchOpen && searchText && i === 0} />
        {/each}
    </div>
    <div class="talon-nav-inner" style="flex: 0 0 auto">
        {#if currentPage.source}
            <a
                class="talon-item"
                href={currentPage.source.url}
                target="_blank"
                referrerpolicy="no-referrer">
                <span class="talon-text">View source</span>
                <Icon
                    iconName={currentPage.source.type}
                    size="40"
                    scale="0.6" />
            </a>
        {/if}
        <div class="talon-item" on:click={openInfo}>
            <span class="talon-text">Info</span>
            <PageIcon page={currentPage} />
        </div>
        <div class="talon-item" on:click={hideSidebar}>
            <span class="talon-text">Hide sidebar</span>
            <Icon iconName="arrowRight" size="40" scale="0.6" />
        </div>
    </div>
</div>

<FloatingButton hide={sidebarShown} on:click={showSidebar} />

<Modals>
    <div slot="backdrop" class="backdrop" on:click={closeModal} />
</Modals>
