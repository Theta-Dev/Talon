<script lang="ts">
    import {selectTextOnFocus} from "../util/inputDirectives"
    // import ClickOutside from "svelte-click-outside"

    import Icon from "./Icon.svelte"
    import ImageIcon from "./ImageIcon.svelte";

    import iconTalon from "../assets/talon.svg";

    enum MenuState {
        Hidden,
        Closed,
        Open,
    }

    function showMenu(): void {
        menuState = MenuState.Closed;
    }

    function openSearch(): void {
        searchOpen = true;
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

    function closeSearch() {
        searchOpen = false
    }

    let menuState = isMobile() ? MenuState.Hidden : MenuState.Closed
    let searchInput: HTMLInputElement
    let searchOpen: boolean = false
</script>

<!--<ClickOutside on:clickoutside={clickOut}>-->
<div class="wrapper">
    <div class="nav-inner">
        <div class="item" class:active={searchOpen} on:click={openSearch}>
            <span class="text"></span>
            <input placeholder="Search..." bind:this={searchInput}
                   on:focusout={closeSearch}
                   use:selectTextOnFocus
            >
            <Icon iconName="search" size="40" scale="0.6"/>
        </div>
        <div class="item" on:click={() => alert("hi")}>
            <span class="text">Talon</span>
            <ImageIcon imageSrc={iconTalon} size="40" scale="0.8"/>
        </div>
        <div class="item">
            <span class="text">Spotify-Gender-Ex</span>
            <ImageIcon
                imageSrc="https://raw.githubusercontent.com/Theta-Dev/Spotify-Gender-Ex/master/assets/logo_square.svg"
                size="40" scale="0.8"/>
        </div>
        <div class="item">
            <span class="text">Test</span>
            <ImageIcon
                imageSrc="" alt="te"
                size="40" scale="0.8"/>
        </div>
    </div>
    <div class="nav-inner">
        <div class="item">
            <span class="text">View source</span>
            <Icon iconName="github" size="40" scale="0.6"/>
        </div>
        <div class="item">
            <span class="text"></span>
            <div class="subitem">
                <ul>
                    <li>
                        <a href="">06.08.2021 16:14 (537eab73)</a>
                    </li>
                    <li>
                        <a href="">09.08.2021 17:46 (546c31b0)</a>
                    </li>
                    <li>
                        <a href="">12.08.2021 21:08 (5a088a48)</a>
                    </li>
                </ul>
            </div>
            <Icon iconName="history" size="40" scale="0.6"/>
        </div>
        <div class="item">
            <span class="text">Info</span>
            <Icon iconName="info" size="40" scale="0.6"/>
        </div>
    </div>
</div>
<!--<div class="fab" class:hide={menuState > MenuState.Hidden}
     on:click={showMenu}>
    <Icon iconName="menu" size="25"/>
</div>-->
<!--</ClickOutside>-->
