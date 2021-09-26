export interface TalonData {
    root_path: string
    current_page: string | null
    current_version: string
    versions: {[key: string]: TalonVersion}
    pages: {[key: string]: TalonPage}
}

export interface TalonPage {
    name: string
    path: string
    color: string
    visibility: TalonVisibility
    image: string | undefined
    source: TalonLink | undefined
}

export enum TalonVisibility {
    FEATURED = "featured",
    SEARCHABLE = "searchable",
    HIDDEN = "hidden",
}

export interface TalonLink {
    url: string
    type: TalonLinkType
}

export enum TalonLinkType {
    LINK = "link",
    GIT = "git",
    GITHUB = "github",
    GITLAB = "gitlab",
    GITEA = "gitea",
    BITBUCKET = "bitbucket",
}

export interface TalonVersion {
    date: string
    name: string | undefined
    user: string
    tags: {[key: string]: string} | undefined
}

export interface Focusable {
    focus(): void
    blur(): void
}

export interface SvelteActionRes {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    update?: (parameters: any) => void
    destroy?: () => void
}
