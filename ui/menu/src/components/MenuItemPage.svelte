<script lang="ts">
    import type {TalonPage} from "../util/types"
    import PageIcon from "./PageIcon.svelte"
    import MenuItem from "./MenuItem.svelte"

    export let page: TalonPage
    export let rootPath: string = "/"
    export let active: boolean = false

    const MAX_TEXT_LEN = 20

    let text: string
    $: text =
        page.name.length > MAX_TEXT_LEN
            ? page.name.substr(0, 20).trim() + "..."
            : page.name

</script>

<MenuItem {active} {text} link={rootPath + page.path}>
    <PageIcon {page} />
</MenuItem>
