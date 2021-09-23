<script lang="ts">
    let hovering

    function enter() {
        hovering = true
    }

    function leave() {
        hovering = false
    }

    export let link: string = null
    export let text: string = ""
    export let active: boolean = false
    export let newTab: boolean = false
    export let privacy: boolean = false

</script>

<style lang="sass">
    @use "../style/values"

    a
        display: flex
        flex-direction: row-reverse
        align-items: center

        text-align: right
        margin-bottom: 0.2em

        border-radius: 20px

        cursor: pointer

        >span
            font-size: 1.1em
            font-weight: bold

            margin-left: auto
            padding: 0 0.5em

            display: none
            pointer-events: none

        &:hover, &.active
            pointer-events: auto
            background-color: values.$color-primary
            text-decoration: none

            >span, >:global(*)
                display: flex
</style>

<a
    class="talon-item"
    class:active={active || hovering}
    on:mouseenter={enter}
    on:mouseleave={leave}
    href={link}
    target={newTab ? '_blank' : null}
    rel={privacy ? 'noopener noreferrer' : null}
    on:click>
    <span>{text}</span>
    <slot />
</a>
