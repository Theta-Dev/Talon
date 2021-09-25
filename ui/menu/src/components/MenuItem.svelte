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
        all: unset

        display: flex
        flex-direction: row-reverse
        align-items: center

        text-align: right
        margin-bottom: 0.2em

        border-radius: 20px

        cursor: pointer

        > span
            font-size: 1.1em
            font-weight: bold

            margin-left: auto
            padding: 0 0.5em

            display: none
            pointer-events: none

        &:hover, &.active
            pointer-events: auto
            background-color: values.$var-primary

            > span, > :global(*)
                display: flex
</style>

<a
    class:active
    href={link}
    target={newTab ? '_blank' : null}
    rel={privacy ? 'noopener noreferrer' : null}
    on:click>
    <span>{text}</span>
    <slot />
</a>
