import type {TalonData, TalonPage, TalonVersion} from "./types"
import {TalonVisibility} from "./types"

const talonData: TalonData = JSON.parse(
    document.getElementById("talon-data").textContent
) as TalonData

const isTalonData = (obj: TalonData) =>
    "root_path" in obj &&
    "current_page" in obj &&
    "current_version" in obj &&
    "versions" in obj &&
    "pages" in obj

const isTalonVersion = (obj: TalonVersion) => "date" in obj && "user" in obj

const rootPath = talonData.root_path

const currentVersion = talonData.versions[talonData.current_version]
const currentVersionId = talonData.current_version

const currentPage: TalonPage = talonData.current_page
    ? talonData.pages[talonData.current_page]
    : {
          name: "#" + talonData.current_version,
          path: "&v/" + talonData.current_version,
          color: "#7935df",
          visibility: TalonVisibility.HIDDEN,
          image: undefined,
          source: undefined,
      }
const currentPageId = talonData.current_page

const isPresent = isTalonData(talonData) && isTalonVersion(currentVersion)

const versions = talonData.versions
const pages = talonData.pages

export {
    rootPath,
    isPresent,
    currentVersion,
    currentVersionId,
    currentPage,
    currentPageId,
    versions,
    pages,
}
