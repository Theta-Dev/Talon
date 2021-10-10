<script lang="ts">
	import {fly} from "svelte/transition"
	import {closeModal} from "svelte-modals"
	import Keydown from "svelte-keydown"

	import type {TalonVersion} from "../util/types"
	import PageIcon from "./PageIcon.svelte"
	import Icon from "./Icon.svelte"
	import {formatDate} from "../util/functions"
	import InlineIcon from "./InlineIcon.svelte"
	import Tag from "./Tag.svelte"
	import {
		currentPage,
		currentVersion,
		currentVersionId,
		rootPath,
		versions,
	} from "../util/talonData"

	export let isOpen: boolean

	function getVersionName(versionId: string, version: TalonVersion): string {
		return version.name ? version.name : "#" + versionId
	}

	function getVersionUrl(versionId: string, version: TalonVersion): string {
		return (
			rootPath +
			(currentPage && version.name
				? currentPage.path + "@" + version.name
				: "&v/" + versionId)
		)
	}

	let versionName: string
	$: versionName = getVersionName(currentVersionId, currentVersion)

	let versionUrl: string
	$: versionUrl = getVersionUrl(currentVersionId, currentVersion)

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
	$: history = Object.entries(versions)
		.filter((e) => e[0] !== currentVersionId)
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

		> div
			position: relative
			overflow-y: auto
			overflow-x: hidden

			margin: 75px auto
			padding: 20px
			width: 600px
			max-width: 90%
			max-height: 90%

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

		&:hover
			filter: brightness(50%)

	.dhead
		width: 100%
		font-size: 1.4em
		border-style: solid
		border-image-source: linear-gradient(to right, values.$color-base-1, values.$color-base-2, values.$color-base-1)
		border-image-slice: 0 0 1 0
		border-image-width: 2px

	.smalltag
		margin: 0.3em 0
		display: flex

		> *
			display: flex

		a
			filter: none

		a:hover
			text-decoration: none
			filter: brightness(130%)

		span
			padding: 0.4em
			background-color: values.$color-base-1

		span:first-child
			font-weight: bold
			background-color: var(--talon-color)

</style>

<Keydown paused={!isOpen} on:Escape={closeModal} />

{#if isOpen}
	<div class="modal" role="dialog" transition:fly={{y: 50}} on:introstart on:outroend>
		<div>
			<div class="tag">
				<PageIcon page={currentPage} size={60} scale={0.8} />
				<span>{currentPage ? currentPage.name : 'v' + currentVersionId}</span>
			</div>

			{#if !currentPage}
				<p>
					This is a dangling version, i.e. it does not belong to a page.
					Assign it to a page or it will be purged within 24 hours.
				</p>
			{/if}

			<p class="dhead">
				<InlineIcon iconName="question" />
				Current version
			</p>

			<Tag key="Version" value={versionName} href={versionUrl} />
			<Tag key="Upload date" value={uploadDate} />
			<Tag key="Uploaded by" value={currentVersion.user} />

			{#each pageTags as [key, value]}
				<Tag {key} {value} />
			{/each}

			{#if history.length}
				<p class="dhead">
					<InlineIcon iconName="history" />
					History
				</p>

				{#each history as [date, name, url]}
					<p class="smalltag">
						<a href={url}> <span>{name}</span> <span>{date}</span> </a>
					</p>
				{/each}
			{/if}

			<p class="dhead" />

			<div>
				This site is powered by
				<a
					href="https://github.com/Theta-Dev/Talon/tree/__VERSION__"
					target="_blank"
					referrerpolicy="no-referrer">Talon __VERSION__</a>, a static site
				management system created by
				<a
					href="https://thetadev.de"
					target="_blank"
					referrerpolicy="no-referrer">ThetaDev</a>
			</div>
			<p><a href={rootPath + '&credits'} target="_blank">View licenses</a></p>
			<button on:click={closeModal}>
				<Icon iconName="close" size={40} scale={0.6} transparent={true} />
			</button>
		</div>
	</div>
{/if}
