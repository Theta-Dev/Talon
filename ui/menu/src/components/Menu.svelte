<script lang="ts">
    import {selectTextOnFocus} from "../util/inputDirectives"
    import Icon from "./Icon.svelte"
    import MenuItem from "./MenuItem.svelte";
    import type {TalonPage} from "../util/types";

    import iconTalon from "../assets/talon.svg";

    function showSidebar(): void {
        sidebarShown = true;
    }

    function hideSidebar(): void {
        sidebarShown = false;
    }

    function isMobile(): boolean {
        return window.innerWidth < 768
    }

    function openSearch(): void {
        searchOpen = true;
        searchInput.focus()
    }

    function closeSearch() {
        searchOpen = false
    }

    const testItem: TalonPage = {
        name: "Talon",
        uri: "",
        color: "green",
        image: iconTalon,
        source: null,
        versions: [],
    }
    const testItem2: TalonPage = {
        name: "Talon, this is just a test",
        uri: "",
        color: "green",
        image: null,
        source: null,
        versions: [],
    }

    let sidebarShown: boolean = !isMobile()
    let searchInput: HTMLInputElement
    let searchOpen: boolean = false
</script>


<div class="wrapper" class:hide={!sidebarShown}>
    <div class="nav-inner" style="flex: 0 0 auto">
        <div class="item" class:active={searchOpen} on:click={openSearch}>
            <span class="text"></span>
            <input placeholder="Search..." bind:this={searchInput}
                   on:focusout={closeSearch}
                   use:selectTextOnFocus
            >
            <Icon iconName="search" size="40" scale="0.6"/>
        </div>
    </div>
    <div class="nav-inner" style="flex: 2 1 auto">
        <MenuItem page={testItem}/>
        <MenuItem page={testItem2}/>
    </div>
    <div class="nav-inner" style="flex: 0 0 auto">
        <div class="item">
            <span class="text">View source</span>
            <Icon iconName="github" size="40" scale="0.6"/>
        </div>
        <div class="item">
            <span class="text">Info</span>
            <Icon iconName="info" size="40" scale="0.6"/>
        </div>
        <div class="item" on:click={hideSidebar}>
            <span class="text">Hide sidebar</span>
            <Icon iconName="arrowRight" size="40" scale="0.6"/>
        </div>
    </div>
</div>
<div class="fab" class:hide={sidebarShown}
     on:click={showSidebar}>
    <Icon iconName="menu" size="25"/>
</div>
