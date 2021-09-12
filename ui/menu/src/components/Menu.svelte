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

    let menuState = isMobile() ? MenuState.Hidden : MenuState.Closed
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
