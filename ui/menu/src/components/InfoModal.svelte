<script lang="ts">
    import {fly} from "svelte/transition"
    import {closeModal} from "svelte-modals"
    import Keydown from "svelte-keydown"

    import type {TalonData, TalonPage, TalonVersion} from "../util/types"
    import PageIcon from "./PageIcon.svelte"
    import Icon from "./Icon.svelte"
    import {formatDate} from "../util/functions";

    export let isOpen: boolean
    export let data: TalonData

    function getVersionName(versionId: string, version: TalonVersion): string {
        return version.name ? version.name : '#' + versionId
    }

    function getVersionUrl(versionId: string, version: TalonVersion): string {
        return data.root_path + (currentPage && version.name ? currentPage.path + '@' + version.name : '&v/' + versionId)
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
            getVersionUrl(key, version)
        ])

</script>

<style lang="sass">
    @use "../style/values"

    talon-modal
        position: fixed
        top: 0
        bottom: 0
        right: 0
        left: 0
        display: flex
        justify-content: center
        align-items: center
        pointer-events: none

        >talon-div
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

        talon-span
            font-size: 2em
            margin-left: 0.25em

    talon-p
        white-space: nowrap

    talon-button
        position: absolute
        right: 5px
        top: 5px

</style>

<Keydown paused={!isOpen} on:Escape={closeModal} />

{#if isOpen}
    <talon-modal
        role="dialog"
        transition:fly={{y: 50}}
        on:introstart
        on:outroend>
        <talon-div>
            <talon-div class="talon-tag">
                <PageIcon page={currentPage} size={60} scale={0.8} />
                <talon-span>{currentPage.name}</talon-span>
            </talon-div>
            <talon-p>Version:
                <a href={versionUrl}>{versionName}</a>
            </talon-p>
            <talon-p>Upload date: {uploadDate}</talon-p>
            <talon-p>Uploaded by: {currentVersion.user}</talon-p>

            {#if currentVersion.tags}
                {#each pageTags as [key, val]}
                    <talon-p>
                        {key}:
                        <talon-code>{val}</talon-code>
                    </talon-p>
                {/each}
            {/if}

            <talon-hr />

            {#each history as [date, name, url]}
                <talon-p>
                    <a href={url}>
                        {date}&nbsp;&nbsp;&nbsp;{name}
                    </a>
                </talon-p>
            {/each}

            <talon-hr />

            <talon-div>
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
            </talon-div>
            <talon-p>
                <a href={data.root_path + 'int/license'} target="_blank">View
                    licenses</a>
            </talon-p>
            <talon-button on:click={closeModal}>
                <Icon
                    iconName="close"
                    size={40}
                    scale={0.6}
                    transparent={true} />
            </talon-button>
        </talon-div>
    </talon-modal>
{/if}
