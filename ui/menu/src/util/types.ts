export interface TalonData {
    root_path: string
    current_page: TalonPage
    pages: TalonPage[]
}

export interface TalonPage {
    name: string
    path: string
    color: string
    image: string | null | undefined
    source: TalonLink | null | undefined
    versions: TalonVersion[]
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
    date: Date
    user: string
    tag: string
}
