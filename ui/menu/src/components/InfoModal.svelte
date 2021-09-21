<script lang="ts">
    import {fly} from "svelte/transition"

    import type {TalonData, TalonVersion} from "../util/types"
    import ImageIcon from "./ImageIcon.svelte"
    // noinspection ES6UnusedImports
    import {version} from "../../package.json"
    import {TalonPage} from "../util/types";

    export let isOpen: boolean
    export let data: TalonData

    let currentPage: TalonPage
    $: currentPage = data.pages[data.current_page]

    let currentVersion: TalonVersion
    $: currentVersion = data.versions[data.current_version]

    let shortName: string
    $: shortName = currentPage.name.substr(0, 2)

    let uploadDate: string
    $: uploadDate = new Date(currentVersion.date).toLocaleString(
        navigator.language
    )

</script>

<style lang="sass">
    @use "../style/values"

    .modal
        position: fixed
        top: 0
        bottom: 0
        right: 0
        left: 0
        display: flex
        justify-content: center
        align-items: center
        pointer-events: none

    .contents
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

    .tag
        display: flex
        align-items: center

        .text
            font-size: 2em
            margin-left: 0.25em

</style>

{#if isOpen}
    <!-- on:introstart and on:outroend are required to transition 1 at a time between modals -->
    <div
        role="dialog"
        class="modal"
        transition:fly={{y: 50}}
        on:introstart
        on:outroend>
        <div class="contents">
            <div class="tag">
                <ImageIcon
                    imageSrc={currentPage.image}
                    color={currentPage.color}
                    alt={shortName}
                    size="40"
                    scale="0.8" />
                <span class="text"> {currentPage.name} </span>
            </div>
            <p>Upload date: {uploadDate}</p>
            <p>Uploaded by: {currentVersion.user}</p>
            <p>Version-ID: <code>{data.current_version}</code></p>

            {#if currentVersion.tags}
                {#each Object.entries(currentVersion.tags) as [key, val]}
                    <p>
                        {key.replace(/^\w/, (c) => c.toUpperCase())}:
                        <code>{val}</code>
                    </p>
                {/each}
            {/if}

            <p>
                This site is powered by
                <a
                    href="https://github.com/Theta-Dev/Talon/tree/{version}"
                    target="_blank"
                    referrerpolicy="no-referrer">Talon
                    {version}</a>, a static site management system created by
                <a
                    href="https://thetadev.de"
                    target="_blank"
                    referrerpolicy="no-referrer">ThetaDev</a>
            </p>
            <p>
                <a href={data.root_path + 'int/license'} target="_blank">View
                    licenses</a>
            </p>
        </div>
    </div>
{/if}
