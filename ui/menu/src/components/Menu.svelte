<script lang="ts">
    import {selectTextOnFocus} from "../util/inputDirectives"
    import ClickOutside from "svelte-click-outside"

    import Icon from "./Icon.svelte"
    import ImageIcon from "./ImageIcon.svelte"

    import talonLogo from "../assets/talon.svg"

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
                <div class="item" on:click={openSearch}>
                    <Icon iconName="search" size="40" scale="0.7"/>
                    <input
                        bind:this={searchInput}
                        use:selectTextOnFocus
                        placeholder="Search"
                    />
                </div>

                <div class="item">
                    <ImageIcon imageSrc={talonLogo} size="40"/>
                    <span class="text">Talon</span>
                </div>

                <div class="item">
                    <ImageIcon
                        imageSrc="https://raw.githubusercontent.com/Theta-Dev/Spotify-Gender-Ex/master/assets/logo_square.svg"
                        size="40"/>
                    <span class="text">Spotify-Gender-Ex</span>
                </div>

                <div class="item">
                    <ImageIcon
                        imageSrc="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fwww.imagesource.com%2Fwp-content%2Fuploads%2F2019%2F06%2FRio.jpg&f=1&nofb=1"
                        size="40"/>
                    <span class="text">Some landscape</span>
                </div>
            </div>
            <div class="nav-inner">
                <div class="item">
                    <Icon iconName="github" size="40" scale="0.7"/>
                    <span class="text">View source</span>
                </div>
                <div class="item" on:click={() => closeMenu(true)}>
                    <Icon iconName={menuIcon} size="40" scale="0.7"/>
                    <span class="text">Close sidebar</span>
                </div>
            </div>
        </nav>
    </div>
    <div class="fab" class:hide={menuState > MenuState.Hidden}
         on:click={showMenu}>
        <Icon iconName="menu" size="25"/>
    </div>
</ClickOutside>
