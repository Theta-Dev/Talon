<script lang="ts">
    import {closeModal, Modals, openModal} from "svelte-modals"
    import {selectTextOnFocus} from "../util/inputDirectives"

    import Icon from "./Icon.svelte"
    import MenuItem from "./MenuItem.svelte"
    import InfoModal from "./InfoModal.svelte"

    import testData from "../testdata/test.json"
    import type {TalonData, TalonPage} from "../util/types"
    import {TalonVisibility} from "../util/types"

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

        if (displayedPages.length === 0) searchText = ""
    }

    function searchKeypress(e: KeyboardEvent) {
        if (e.key === "Enter" && displayedPages) {
            window.location = talonData.root_path + displayedPages[0].path
        }
    }

    function openInfo() {
        openModal(InfoModal, {
            data: talonData,
        })
    }

    const talonData: TalonData = testData

    let sidebarShown: boolean = !isMobile()
    let searchInput: HTMLInputElement
    let searchOpen: boolean = false
    let searchText: string = ""

    let currentPage: TalonPage
    $: currentPage = talonData.pages[talonData.current_page]

    let displayedPages: TalonPage[]
    $: displayedPages = Object.values(talonData.pages).filter((page) => {
        if (searchText) {
            return (
                page.visibility !== TalonVisibility.HIDDEN &&
                page.name.toLowerCase().includes(searchText.toLowerCase())
            )
        }
        return page.visibility === TalonVisibility.FEATURED
    })

</script>

<style lang="sass">
    .backdrop
        position: fixed
        top: 0
        bottom: 0
        right: 0
        left: 0
        background: rgba(0, 0, 0, 0.6)
</style>

<div class="wrapper" class:hide={!sidebarShown}>
    <div class="nav-inner" style="flex: 0 0 auto">
        <div class="item" class:active={searchOpen} on:click={openSearch}>
            <span class="text" />
            <input
                placeholder="Search..."
                bind:this={searchInput}
                bind:value={searchText}
                on:focusout={closeSearch}
                on:keypress={searchKeypress}
                use:selectTextOnFocus />
            <Icon iconName="search" size="40" scale="0.6" dot={searchText} />
        </div>
    </div>
    <div class="nav-inner" style="flex: 2 1 auto">
        {#each displayedPages as page, i}
            <MenuItem
                {page}
                rootPath={talonData.root_path}
                active={searchOpen && searchText && i === 0} />
        {/each}
    </div>
    <div class="nav-inner" style="flex: 0 0 auto">
        {#if currentPage.source}
            <a
                class="item"
                href={currentPage.source.url}
                target="_blank"
                referrerpolicy="no-referrer">
                <span class="text">View source</span>
                <Icon
                    iconName={currentPage.source.type}
                    size="40"
                    scale="0.6" />
            </a>
        {/if}
        <div class="item" on:click={openInfo}>
            <span class="text">Info</span>
            <Icon iconName="info" size="40" scale="0.6" />
        </div>
        <div class="item" on:click={hideSidebar}>
            <span class="text">Hide sidebar</span>
            <Icon iconName="arrowRight" size="40" scale="0.6" />
        </div>
    </div>
</div>
<div class="fab" class:hide={sidebarShown} on:click={showSidebar}>
    <Icon iconName="menu" size="25" />
</div>

<Modals>
    <div slot="backdrop" class="backdrop" on:click={closeModal} />
</Modals>
