<script lang="ts">
    import ImageIcon from "./ImageIcon.svelte"
    import type {TalonPage} from "../util/types"

    export let page: TalonPage
    export let size: number = 40
    export let scale: number = 0.8

    let shortName: string
    $: shortName = page.name.substr(0, 2)

</script>

<ImageIcon
    imageSrc={page.image}
    color={page.color}
    alt={shortName}
    {size}
    {scale} />
