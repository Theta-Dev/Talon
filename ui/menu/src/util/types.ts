export interface TalonPage {
    name: string
    uri: string
    color: string
    image: string | null
    source: TalonLink | null
    versions: TalonVersion[]
}

export interface TalonLink {
    url: string
    type: TalonLinkType
}

export enum TalonLinkType {
    GENERIC = "generic",
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
