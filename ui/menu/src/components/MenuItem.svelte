<script lang="ts">
    export let link: string = null
    export let text = ""
    export let active = false
    export let newTab = false
    export let privacy = false

</script>

<style lang="sass">
    @use "../style/values"

    a
        all: unset !important

        display: flex !important
        flex-direction: row-reverse !important
        align-items: center !important

        text-align: right !important
        margin-bottom: 0.2em !important

        border-radius: 20px !important

        cursor: pointer !important

        >talon-span
            font-size: 1.1em
            font-weight: bold

            margin-left: auto
            padding: 0 0.5em

            display: none
            pointer-events: none

        &:hover, &.active
            pointer-events: auto !important
            background-color: values.$var-primary !important

            >talon-span, >:global(*)
                display: flex !important
</style>

<a
    class:active
    href={link}
    target={newTab ? '_blank' : null}
    rel={privacy ? 'noopener noreferrer' : null}
    on:click>
    <talon-span>{text}</talon-span>
    <slot />
</a>
