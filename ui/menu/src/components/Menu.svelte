<script lang="ts">
    import {selectTextOnFocus} from "../util/inputDirectives"
    import Icon from "./Icon.svelte"
    import MenuItem from "./MenuItem.svelte"

    import testData from "../testdata/test.json"
    import type {TalonData} from "../util/types"

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
    }

    const talonData: TalonData = testData

    let sidebarShown: boolean = !isMobile()
    let searchInput: HTMLInputElement
    let searchOpen: boolean = false

</script>

<div class="wrapper" class:hide={!sidebarShown}>
    <div class="nav-inner" style="flex: 0 0 auto">
        <div class="item" class:active={searchOpen} on:click={openSearch}>
            <span class="text" />
            <input
                placeholder="Search..."
                bind:this={searchInput}
                on:focusout={closeSearch}
                use:selectTextOnFocus />
            <Icon iconName="search" size="40" scale="0.6" />
        </div>
    </div>
    <div class="nav-inner" style="flex: 2 1 auto">
        {#each talonData.pages as page}
            <MenuItem {page} rootPath={talonData.root_path} />
        {/each}
    </div>
    <div class="nav-inner" style="flex: 0 0 auto">
        {#if talonData.current_page.source}
            <a
                class="item"
                href={talonData.current_page.source.url}
                target="_blank"
                referrerpolicy="no-referrer">
                <span class="text">View source</span>
                <Icon
                    iconName={talonData.current_page.source.type}
                    size="40"
                    scale="0.6" />
            </a>
        {/if}
        <div class="item">
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
