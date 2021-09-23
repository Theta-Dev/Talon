<script lang="ts">
    import {fly} from "svelte/transition"
    import {closeModal} from "svelte-modals"
    import Keydown from "svelte-keydown"

    import type {TalonData, TalonPage, TalonVersion} from "../util/types"
    import PageIcon from "./PageIcon.svelte"
    import Icon from "./Icon.svelte"

    export let isOpen: boolean
    export let data: TalonData

    let currentPage: TalonPage
    $: currentPage = data.pages[data.current_page]

    let currentVersion: TalonVersion
    $: currentVersion = data.versions[data.current_version]

    let uploadDate: string
    $: uploadDate = new Date(currentVersion.date).toLocaleString(
        /* global navigator */
        navigator.language
    )

    let pageTags: [string, string][]
    $: pageTags = currentVersion.tags
        ? Object.entries(currentVersion.tags).map(([key, val]) => [
              key.replace(/^\w/, (c) => c.toUpperCase()),
              val,
          ])
        : []

</script>

<style lang="sass">
    @use "../style/values"

    .talon-modal
        position: fixed
        top: 0
        bottom: 0
        right: 0
        left: 0
        display: flex
        justify-content: center
        align-items: center
        pointer-events: none

    .talon-contents
        position: relative
        overflow: auto

        margin: 75px auto
        padding: 20px
        width: 600px
        max-width: 90%

        background: values.$color-base
        border-radius: 15px
        box-shadow: 0 0 50px rgba(0, 0, 0, 0.5)

        pointer-events: auto

    .talon-tag
        display: flex
        align-items: center

        .talon-text
            font-size: 2em
            margin-left: 0.25em

    .talon-close
        position: absolute
        right: 5px
        top: 5px

</style>

<Keydown paused={!isOpen} on:Escape={closeModal} />

{#if isOpen}
    <div
        role="dialog"
        class="talon-modal"
        transition:fly={{y: 50}}
        on:introstart
        on:outroend>
        <div class="talon-contents">
            <div class="talon-tag">
                <PageIcon page={currentPage} size={60} scale={0.8} />
                <span class="talon-text"> {currentPage.name} </span>
            </div>
            <p>Upload date: {uploadDate}</p>
            <p>Uploaded by: {currentVersion.user}</p>

            {#if currentVersion.tags}
                {#each pageTags as [key, val]}
                    <p>{key}:<code>{val}</code></p>
                {/each}
            {/if}

            <hr />

            <p>
                This site is powered by
                <a
                    href="https://github.com/Theta-Dev/Talon/tree/{data.talon_version}"
                    target="_blank"
                    referrerpolicy="no-referrer">Talon
                    {data.talon_version}</a>, a static site management system
                created by
                <a
                    href="https://thetadev.de"
                    target="_blank"
                    referrerpolicy="no-referrer">ThetaDev</a>
            </p>
            <p>
                <a href={data.root_path + 'int/license'} target="_blank">View
                    licenses</a>
            </p>
            <button class="talon-close" on:click={closeModal}>
                <Icon iconName="close" size={40} scale={0.6} />
            </button>
        </div>
    </div>
{/if}
