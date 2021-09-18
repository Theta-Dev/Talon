<script lang="ts">
    import ImageIcon from "./ImageIcon.svelte"
    import type { TalonPage } from "../util/types"

    export let page: TalonPage

    const MAX_TEXT_LEN = 20

    let text: string
    $: text =
        page.name.length > MAX_TEXT_LEN
            ? page.name.substr(0, 20).trim() + "..."
            : page.name

    let shortName: string
    $: shortName = page.name.substr(0, 2)

</script>

<div class="item">
    <span class="text">{text}</span>
    <ImageIcon
        imageSrc={page.image}
        color={page.color}
        alt={shortName}
        size="40"
        scale="0.8" />
</div>
