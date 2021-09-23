export interface TalonData {
    talon_version: string
    root_path: string
    current_page: string
    current_version: string
    versions: {[key: string]: TalonVersion}
    pages: {[key: string]: TalonPage}
}

export interface TalonPage {
    name: string
    path: string
    color: string
    visibility: TalonVisibility
    image: string | null | undefined
    source: TalonLink | null | undefined
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
    id: number
    date: Date
    user: string
    tags: object
}

export interface Focusable {
    focus(): void
    blur(): void
}
