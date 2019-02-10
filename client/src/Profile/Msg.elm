module Profile.Msg exposing (Msg(..), idxToTab, tabToIdx)

import Http
import Material
import Profile.Model exposing (Profile, UserAssets)
import Profile.Ports exposing (ImagePortData)
import Profile.Types exposing (ViewTab(..))
import Timeline.Msg


type Msg
    = Mdl (Material.Msg Msg)
    | SwitchView ViewTab
    | OnProfileLoadResponse (Result Http.Error Profile)
    | UploadProfileImage
    | OnUploadProfileImageResponse (Result Http.Error {})
    | ImageSelected
    | LoadMore
    | ImageRead ImagePortData
    | TimelineMsg Timeline.Msg.Msg
    | OnUserAssetsLoadResponse (Result Http.Error UserAssets)


idxToTab : Int -> Msg
idxToTab idx =
    case idx of
        0 ->
            SwitchView TimelineTab

        1 ->
            SwitchView TokensTab

        2 ->
            SwitchView WalletTab

        _ ->
            SwitchView TimelineTab


tabToIdx : ViewTab -> Int
tabToIdx tab =
    case tab of
        TimelineTab ->
            0

        TokensTab ->
            1

        WalletTab ->
            2
