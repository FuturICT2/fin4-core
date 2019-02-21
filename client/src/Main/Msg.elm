module Main.Msg exposing (Msg(..))

import Asset.Msg
import Common.Json exposing (EmptyResponse)
import CreateAsset.Msg
import ExploreAssets.Msg
import Homepage.Msg
import Http
import Main.User exposing (User)
import Material
import Navigation exposing (Location)
import Portfolio.Msg
import Profile.Msg
import UserLogin.Msg
import Window


type Msg
    = OnRouteChange Location
    | ToggleMobileNav
    | OnCheckSessionResponse (Result Http.Error User)
    | Mdl (Material.Msg Msg)
    | OnWindowResize Window.Size
    | HomepageMsg Homepage.Msg.Msg
    | ExploreAssetsMsg ExploreAssets.Msg.Msg
    | Portfolio Portfolio.Msg.Msg
    | UserLoginMsg UserLogin.Msg.Msg
    | UserLogout
    | OnLogoutResponse (Result Http.Error EmptyResponse)
    | CreateAssetMsg CreateAsset.Msg.Msg
    | AssetMsg Asset.Msg.Msg
    | ProfileMsg Profile.Msg.Msg
