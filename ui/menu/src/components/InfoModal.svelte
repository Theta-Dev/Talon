<script lang="ts">
    import {fly} from "svelte/transition"
    import {closeModal} from "svelte-modals"
    import Keydown from "svelte-keydown"

    import type {TalonData, TalonPage, TalonVersion} from "../util/types"
    import PageIcon from "./PageIcon.svelte"
    import Icon from "./Icon.svelte"
    import {formatDate} from "../util/functions"

    export let isOpen: boolean
    export let data: TalonData

    function getVersionName(versionId: string, version: TalonVersion): string {
        return version.name ? version.name : "#" + versionId
    }

    function getVersionUrl(versionId: string, version: TalonVersion): string {
        return (
            data.root_path +
            (currentPage && version.name
                ? currentPage.path + "@" + version.name
                : "&v/" + versionId)
        )
    }

    let currentPage: TalonPage
    $: currentPage = data.pages[data.current_page]

    let currentVersion: TalonVersion
    $: currentVersion = data.versions[data.current_version]

    let versionName: string
    $: versionName = getVersionName(data.current_version, currentVersion)

    let versionUrl: string
    $: versionUrl = getVersionUrl(data.current_version, currentVersion)

    let uploadDate: string
    $: uploadDate = formatDate(currentVersion.date)

    let pageTags: [string, string][]
    $: pageTags = currentVersion.tags
        ? Object.entries(currentVersion.tags).map(([key, val]) => [
              key.replace(/^\w/, (c) => c.toUpperCase()),
              val,
          ])
        : []

    let history: [string, string, string][]
    $: history = Object.entries(data.versions)
        .filter((e) => e[0] !== data.current_version)
        .map(([key, version]) => [
            formatDate(version.date),
            getVersionName(key, version),
            getVersionUrl(key, version),
        ])

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

        >div
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

        span
            font-size: 2em
            margin-left: 0.25em

    button
        position: absolute
        right: 5px
        top: 5px
        background: none
        border: none

    p
        margin: 0.8em 0
        white-space: nowrap

    code
        font-family: monospace
        background-color: values.$color-base-1
        border-radius: 0.3em
        padding: 0.1em 0.3em

    hr
        width: 100%
        height: 2px
        margin: 1.6em 0
        border: none
        background-image: linear-gradient(to right, values.$color-base-1, values.$color-base-2, values.$color-base-1)

    a
        display: inline
        color: values.$color-primary-light
        text-decoration: none

        &:hover
            text-decoration: underline

</style>

<Keydown paused={!isOpen} on:Escape={closeModal} />

{#if isOpen}
    <div
        class="modal"
        role="dialog"
        transition:fly={{y: 50}}
        on:introstart
        on:outroend>
        <div>
            <div class="tag">
                <PageIcon page={currentPage} size={60} scale={0.8} />
                <span>{currentPage.name}</span>
            </div>
            <p>Version: <a href={versionUrl}>{versionName}</a></p>
            <p>Upload date: {uploadDate}</p>
            <p>Uploaded by: {currentVersion.user}</p>

            {#if currentVersion.tags}
                {#each pageTags as [key, val]}
                    <p>{key}: <code>{val}</code></p>
                {/each}
            {/if}

            <hr />

            {#each history as [date, name, url]}
                <p><a href={url}> {date}&nbsp;&nbsp;&nbsp;{name} </a></p>
            {/each}

            <hr />

            <div>
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
            </div>
            <p>
                <a href={data.root_path + 'int/license'} target="_blank">View
                    licenses</a>
            </p>
            <button on:click={closeModal}>
                <Icon
                    iconName="close"
                    size={40}
                    scale={0.6}
                    transparent={true} />
            </button>
        </div>
    </div>
{/if}
