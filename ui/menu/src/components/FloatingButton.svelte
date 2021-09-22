<script lang="ts">
    import Icon from "./Icon.svelte"

    export let hide: boolean = false

</script>

<style lang="sass">
    @use "../style/values"

    .talon-fab
        position: fixed
        bottom: 25px
        right: 25px

        height: 56px
        width: 56px

        display: flex
        align-items: center
        justify-content: center

        border-radius: 50%

        background: values.$color-primary
        opacity: 1

        transition: bottom 0.5s, opacity 0.5s

        &.talon-hide
            bottom: -56px
            opacity: 0
</style>

<button class="talon-fab" class:talon-hide={hide} on:click>
    <Icon iconName="menu" size="25" />
</button>
